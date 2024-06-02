package commands

import "github.com/charmbracelet/log"

type UninstallCommand struct {
	Port string `arg:""`
}

func (i *UninstallCommand) Run(ctx *Context) error {
	log.Fatal("Not implemented")
	return nil
}
