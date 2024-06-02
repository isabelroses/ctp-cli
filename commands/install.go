package commands

import "github.com/charmbracelet/log"

type InstallCommand struct {
	Port string `arg:""`
}

func (i *InstallCommand) Run(ctx *Context) error {
	log.Fatal("Not implemented")
	return nil
}
