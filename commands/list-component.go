package commands

import (
	"fmt"
	"io"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	catppuccin "github.com/catppuccin/go"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(catppuccin.Mocha.Mauve().Hex))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type RenderableItem interface {
	Render() string
}

type ListComponent[TItem RenderableItem] struct {
	list list.Model
}

func (m ListComponent[TItem]) Init() tea.Cmd {
	return nil
}

func (m ListComponent[TItem]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ListComponent[TItem]) View() string {
	return "\n" + m.list.View()
}

func (l ListComponent[TItem]) Show() {
	if _, err := tea.NewProgram(l).Run(); err != nil {
		log.Error("Error running program:", err)
	}
}

type ListItemDelegate[TItem RenderableItem] struct {}

func (d ListItemDelegate[TItem]) Height() int                             { return 1 }
func (d ListItemDelegate[TItem]) Spacing() int                            { return 0 }
func (d ListItemDelegate[TItem]) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d ListItemDelegate[TItem]) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(TItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Render())

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type DelegatedItem[TItem RenderableItem] struct { item TItem }
func (d DelegatedItem[TItem]) FilterValue() string {
	return d.item.Render()
}

func NewListComponent[TItem RenderableItem](title string, items []TItem) ListComponent[TItem] {
	const defaultWidth = 20

	delegate := ListItemDelegate[TItem]{}

	mapped := make([]list.Item, 0, len(items))
	for _, v := range items {
		mapped = append(mapped, DelegatedItem[TItem]{ item: v })
	}

	lst := list.New(mapped, delegate, defaultWidth, listHeight)
	lst.Title = title
	lst.SetShowStatusBar(true)
	lst.SetFilteringEnabled(true)
	lst.ShowFilter()
	lst.Styles.Title = titleStyle
	lst.Styles.PaginationStyle = paginationStyle
	lst.Styles.HelpStyle = helpStyle
	lst.Styles.StatusBar = itemStyle


	return ListComponent[TItem]{
		list: lst,
	}
}
