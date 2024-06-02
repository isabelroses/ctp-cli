package main

import (
	"github.com/alecthomas/kong"
	"github.com/catppuccin/cli/commands"
	"github.com/charmbracelet/log"
)

var cli struct {
	Init        commands.InitCommand `cmd:"" help:"Initialise a port from catppuccin/template" aliases:"innit"`
	List        commands.ListCommand `cmd:"" help:"Query ports and userstyles"`
	Interactive bool                 `help:"Enable interactive mode"`
}

func main() {
	log.SetReportTimestamp(false)

	ctx := kong.Parse(&cli,
		kong.UsageOnError(),
		kong.Name("ctp"),
		kong.Description("A suite of tools to help you create and manage our ports"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	err := ctx.Run(&commands.Context{
		Interactive: cli.Interactive,
	})

	ctx.FatalIfErrorf(err)
}
