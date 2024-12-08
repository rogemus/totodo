package cmd

import (
	"fmt"
	"strconv"
	repo "totodo/pkg/repository"
)

type startCmd struct {
	Cmd  string
	repo repo.TasksRepository
}

func NewStartCmd(repo repo.TasksRepository) startCmd {
	return startCmd{
		repo: repo,
		Cmd:  "start",
	}
}

func (cmd startCmd) Run(args []string) {
	if len(args) == 0 {
		fmt.Println("no report type selected")
		return
	}

	// TODO: handle error
	id, _ := strconv.Atoi(args[0])
	task, err := cmd.repo.GetTask(id)

	if err != nil {
		fmt.Printf("no task with id: %d", id)
	}

	fmt.Printf("starting working on task: (#%d) %s", task.Id, task.Description)
}

func (cmd startCmd) Help() {
	fmt.Println("start - help")
}
