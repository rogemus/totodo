package cmd

import (
	"fmt"
	"totodo/pkg"
)

type editCmd struct {
	Cmd  string
	repo pkg.TasksRepository
}

func NewEditCmd(repo pkg.TasksRepository) editCmd {
	return editCmd{
		repo: repo,
		Cmd:  "edit",
	}
}

func (cmd editCmd) Run(args []string) {
	if len(args) == 0 {
		fmt.Println("no report type selected")
		return
	}

	fmt.Println("edit")
}

func (cmd editCmd) Help() {
	fmt.Println("edit - help")
}
