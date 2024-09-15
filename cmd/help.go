package cmd

import (
	"fmt"
	"totodo/pkg"
)

type helpCmd struct {
	commands []pkg.Cmd
	Cmd      string
}

func NewHelpCmd(commands []pkg.Cmd) helpCmd {
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
