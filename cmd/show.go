package cmd

import (
	"fmt"
	"os"
	"strconv"
	"text/template"
	repo "totodo/pkg/repository"
)

type showCmd struct {
	Cmd  string
	repo repo.TasksRepository
}

func NewShowCmd(repo repo.TasksRepository) showCmd {
	return showCmd{
		repo: repo,
		Cmd:  "show",
	}
}

func (cmd showCmd) Run(args []string) {
	if len(args) == 0 {
		fmt.Println("no task id provided")
		return
	}

	// TODO: handle error
	id, _ := strconv.Atoi(args[0])

	// TODO use lipgoss for proper template
	tmpl :=
		`
      id  |  description  |  created
    ----------------------------------
      {{ .Id }}  |  {{ .Description }}  |  {{ .Created }}
    `
	task, err := cmd.repo.GetTask(id)

	if err != nil {
		fmt.Printf("no task with id: %d", id)
	}

	t := template.Must(template.New("list").Parse(tmpl))

	if err := t.Execute(os.Stdout, task); err != nil {
		fmt.Printf("%v", err)
	}
}

func (cmd showCmd) Help() {
	fmt.Println("show -help")
}
