package commands

import (
	"fmt"

	"github.com/catppuccin/cli/query"
	"github.com/catppuccin/cli/shared"
	"github.com/charmbracelet/log"
	"golang.org/x/exp/slices"
)

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

func (r *repositorySubCommand) Run(ctx *shared.Context) error {
	log.Fatal("Not implemented")
	return nil
}

func (r *userstylesSubCommand) Run(ctx *shared.Context) error {
	log.Fatal("Not implemented")
	return nil
}

func (r *portsSubCommand) Run(ctx *shared.Context) error {
	ports := ctx.Datastore.Ports()
	title := "All ports"
	if r.Maintained {
		ports = query.FilterMaintained(ports)
		title = "Maintained ports"
	}

	if r.Unmaintained {
		ports = query.FilterUnmaintained(ports)
		title = "Unmaintained ports"
	}

	if ctx.Interactive {
		showPortList(ports, title)
		return nil
	}

	fmt.Printf("There are %d ", len(ports))
	if r.Maintained {
		fmt.Println("maintained ports")
	}

	if r.Unmaintained {
		fmt.Println("unmaintained ports")
	}

	fmt.Println()

	slices.SortFunc(ports, query.SortPortLexicographically)

	for _, port := range ports {
		fmt.Printf("%s has %d maintainers\n", port.Name, len(port.CurrentMaintainers))
	}

	return nil
}
