package commands

import (
	"github.com/catppuccin/cli/shared"
	"github.com/charmbracelet/log"
)

type UninstallCommand struct {
	Port string `arg:""`
}

func (i *UninstallCommand) Run(ctx *shared.Context) error {
	log.Fatal("Not implemented")
	return nil
}
