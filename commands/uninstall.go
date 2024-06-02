package commands

import (
	"github.com/charmbracelet/log"
)

import "github.com/catppuccin/cli/shared"

type UninstallCommand struct {
	Port string `arg:""`
}

func (i *UninstallCommand) Run(ctx *shared.Context) error {
	log.Fatal("Not implemented")
	return nil
}
