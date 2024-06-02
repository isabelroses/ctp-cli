package commands

import "github.com/charmbracelet/log"

type flags struct {
	Maintained   bool `help:"List only the maintained ports" group:"X" xor:"X"`
	Unmaintained bool `help:"List only the unmaintained ports" group:"X" xor:"X"`
	Archived     bool `help:"List only the archived ports" group:"X" xor:"X"`
}

type ListCommand struct {
	Repositories repositorySubCommand `cmd:"List both ports and userstyles" default:"1" aliases:"repos"`
	Userstyles   userstylesSubCommand `cmd:"List only the userstyles"`
	Ports        portsSubCommand      `cmd:"List only the ports"`
}

type repositorySubCommand struct {
	flags
}

type userstylesSubCommand struct {
	flags
}

type portsSubCommand struct {
	flags
}

func (r *repositorySubCommand) Run(ctx *Context) error {
	log.Fatal("Not implemented")
	return nil
}
