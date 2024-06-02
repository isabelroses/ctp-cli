package commands

import (
	"github.com/catppuccin/cli/shared"
	"github.com/charmbracelet/log"
)

type InstallCommand struct {
	Port string `arg:""`
}

func (i *InstallCommand) Run(ctx *shared.Context) error {
	log.Fatal("Not implemented")
	return nil
}
