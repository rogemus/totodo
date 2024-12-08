package cmd

import (
	"fmt"
	"totodo/pkg/model"
)

type helpCmd struct {
	commands []model.Cmd
	Cmd      string
}

func NewHelpCmd(commands []model.Cmd) helpCmd {
	return helpCmd{
		Cmd:      "help",
		commands: commands,
	}
}

func (cmd helpCmd) Run() {
	fmt.Printf("Help...")

	for _, cmd := range cmd.commands {
		cmd.Help()
	}
}
