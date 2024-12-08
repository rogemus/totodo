package cmd

import (
	"fmt"
	"os"
	"text/template"
	repo "totodo/pkg/repository"
)

type reportCmd struct {
	Cmd  string
	repo repo.TasksRepository
}

func NewReportCmd(repo repo.TasksRepository) reportCmd {
	return reportCmd{
		repo: repo,
		Cmd:  "report",
	}
}

func (cmd reportCmd) Run(args []string) {
	// TODO display available reports
	if len(args) == 0 {
		fmt.Println("no report type selected")
		return
	}

	reportType := args[0]

	// TODO add different report types
	switch reportType {
	case "list":
		// TODO use lipgoss for proper template
		tmpl :=
			`
      id  |  description  |  created
    ----------------------------------
      {{ range . }}{{ .Id }}  |  {{ .Description }}  |  {{ .Created }}
      {{ end }}
    `
		tasks, _ := cmd.repo.GetTasks()
		t := template.Must(template.New("list").Parse(tmpl))

		if err := t.Execute(os.Stdout, tasks); err != nil {
			fmt.Printf("%v", err)
		}
	}
}

func (cmd reportCmd) Help() {
	fmt.Println("repot - help")
}
