package main

import (
	"github.com/alecthomas/kong"
	"github.com/catppuccin/cli/commands"
	"github.com/catppuccin/cli/query"
	"github.com/catppuccin/cli/shared"
	catppuccin "github.com/catppuccin/go"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var cli struct {
	Init        commands.InitCommand      `cmd:"" help:"Initialise a port from catppuccin/template" aliases:"innit"`
	List        commands.ListCommand      `cmd:"" help:"Query ports and userstyles"`
	Install     commands.InstallCommand   `cmd:"" help:"Install a port"`
	Uninstall   commands.UninstallCommand `cmd:"" help:"Uninstall a port"`
	Interactive bool                      `help:"Enable interactive mode"`
}

func main() {
	log.SetReportTimestamp(false)
	log.SetStyles(setLogColours())

	ctx := kong.Parse(&cli,
		kong.UsageOnError(),
		kong.Name("ctp"),
		kong.Description("Soothing pastel theme manager for Catppuccin!"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Tree: true,
		}),
	)

	err := ctx.Run(&shared.Context{
		Interactive: cli.Interactive,
		Datastore:   query.NewDatastore(),
	})

	ctx.FatalIfErrorf(err)
}

func setLogColours() *log.Styles {
	styles := log.DefaultStyles()
	flavour := catppuccin.Mocha
	dot := "•"
	cross := "❌"
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString(dot).
		PaddingLeft(1).
		Foreground(lipgloss.Color(flavour.Teal().Hex))
	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString(dot).
		PaddingLeft(1).
		Foreground(lipgloss.Color(flavour.Subtext1().Hex))
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString(dot).
		PaddingLeft(1).
		Foreground(lipgloss.Color(flavour.Yellow().Hex))
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString(cross).
		PaddingLeft(1).
		Foreground(lipgloss.Color(flavour.Red().Hex))
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString(cross).
		PaddingLeft(1).
		Foreground(lipgloss.Color(flavour.Red().Hex))
	return styles
}
