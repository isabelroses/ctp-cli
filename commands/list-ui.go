package commands

import (
	"fmt"
	"io"
	"strings"

	"github.com/catppuccin/cli/query"
	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const listHeight = 14

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color(catppuccin.Mocha.Mauve().Hex))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type portListItem query.Port

func (i portListItem) FilterValue() string { return i.Name }

type portListItemDelegate struct{}

func (d portListItemDelegate) Height() int                             { return 1 }
func (d portListItemDelegate) Spacing() int                            { return 0 }
func (d portListItemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d portListItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(portListItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

type portListModel struct {
	list     list.Model
	quitting bool
}

func (m portListModel) Init() tea.Cmd {
	return nil
}

func (m portListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m portListModel) View() string {
	return "\n" + m.list.View()
}

func showPortList(ports []query.Port, title string) {
	const defaultWidth = 20

	items := make([]list.Item, 0, len(ports))
	for _, v := range ports {
		items = append(items, portListItem(v))
	}

	l := list.New(items, portListItemDelegate{}, defaultWidth, listHeight)
	l.Title = title
	l.SetShowStatusBar(true)
	l.SetFilteringEnabled(true)
	l.ShowFilter()
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle
	l.Styles.StatusBar = itemStyle

	m := portListModel{list: l}

	if _, err := tea.NewProgram(m).Run(); err != nil {
		log.Error("Error running program:", err)
	}
}
