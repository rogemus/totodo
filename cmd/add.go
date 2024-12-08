package cmd

import (
	"flag"
	"fmt"
	"totodo/pkg/model"
	repo "totodo/pkg/repository"
)

type addCmd struct {
	Cmd  string
	repo repo.TasksRepository
}

func NewAddCmd(repo repo.TasksRepository) addCmd {
	return addCmd{
		repo: repo,
		Cmd:  "add",
	}
}

func (cmd addCmd) Run(args []string) {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)

	// Task description
	desc := fs.String("desc", "", "`description` of the task")
	fs.StringVar(desc, "d", *desc, "alias for -desc")

	// Task tags
	// TODO: support multiple tags
	tag := fs.String("tag", "", "`tags` of the task")
	fs.StringVar(tag, "t", *tag, "alias for -tag")

	// Task project
	proj := fs.String("proj", "", "`project` of the task")
	fs.StringVar(proj, "p", *proj, "alias for -p")

	if err := fs.Parse(args); err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	task := model.NewTask(*desc)
	cmd.repo.CreateTask(task)

	fmt.Printf("added task: %s, +%s, @%s", *desc, *tag, *proj)
}

func (cmd addCmd) Help() {
	fmt.Println("add - help")
}
