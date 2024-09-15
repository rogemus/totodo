package cmd

import (
	"fmt"
	"strconv"
	"totodo/pkg"
)

type stopCmd struct {
	Cmd  string
	repo pkg.TasksRepository
}

func NewStopCmd(repo pkg.TasksRepository) stopCmd {
	return stopCmd{
		repo: repo,
		Cmd:  "stop",
	}
}

func (cmd stopCmd) Run(args []string) {
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

	// TODO: check if task is in progress
	fmt.Printf("stoping working on task: (#%d) %s", task.Id, task.Description)
}

func (cmd stopCmd) Help() {
	fmt.Println("stop - help")
}
