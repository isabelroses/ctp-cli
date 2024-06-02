package main

import (
	"github.com/catppuccin/cli/commands"
	"github.com/alecthomas/kong"
)

var cli struct {
  Init commands.InitCommand `cmd:"" help:"Initialise a port" aliases:"innit"`
  Interactive bool `help:"Enable interactive mode"`
}

func main() {
	ctx := kong.Parse(&cli, 
		kong.UsageOnError(),
		kong.Name("ctp"),
		kong.Description("A suite of tools to help you create and manage our ports"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)

	err := ctx.Run(&commands.Context {
		Interactive: cli.Interactive,
	})

	ctx.FatalIfErrorf(err)
}
