package commands

import (
	"github.com/charmbracelet/log"
)
import "github.com/catppuccin/cli/shared"

type InstallCommand struct {
	Port string `arg:""`
}

func (i *InstallCommand) Run(ctx *shared.Context) error {
	log.Fatal("Not implemented")
	return nil
}
