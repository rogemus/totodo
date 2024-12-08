package cmd

import (
	"flag"
	"fmt"
	repo "totodo/pkg/repository"
)

type editCmd struct {
	Cmd  string
	repo repo.TasksRepository
}

func NewEditCmd(repo repo.TasksRepository) editCmd {
	return editCmd{
		repo: repo,
		Cmd:  "edit",
	}
}

func (cmd editCmd) Run(args []string) {
	if len(args) == 0 {
		fmt.Println("no args")
		return
	}

	// TODO if not args
	fs := flag.NewFlagSet("edit", flag.ContinueOnError)
	id := fs.Int("id", 0, "`id` of the task")
	fs.IntVar(id, "i", *id, "alias for -id")

	// Task description
	desc := fs.String("desc", "", "`description` of the task")
	fs.StringVar(desc, "d", *desc, "alias for -desc")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("error: %s", err)
		return
	}


	task, _ := cmd.repo.GetTask(*id)

	if *desc != "" {
		task.Description = *desc
	}

	cmd.repo.UpdateTask(task)

	fmt.Println(task.Description)
	fmt.Printf("edited task: %d, %s", *id, *desc)
}

func (cmd editCmd) Help() {
	fmt.Println("edit - help")
}
