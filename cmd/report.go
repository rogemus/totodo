package cmd

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	repo "totodo/pkg/repository"
)

type reportCmd struct {
	Cmd  string
	repo repo.TasksRepository
	fs   *flag.FlagSet
}

var reportFlagValues = struct {
	reportTypeFlag string
	listTypesFlag  bool
	helpFlag       bool
}{
	reportTypeFlag: "list",
	listTypesFlag:  false,
	helpFlag:       false,
}

const reportUsageLong = `
  Report tasks:
    -t, --type         typo of the report
    -lt, --list-types  list available reports type
    -h, --help prints  help information
`
const reportUsageShort = `
  Report task:
    -t, --type         typo of the report
    -lt, --list-types  list available reports type
    -d, --desc         description of the task
`

func NewReportCmd(repo repo.TasksRepository) reportCmd {
	cmd := "report"

	set := flag.NewFlagSet(cmd, flag.ContinueOnError)
	set.StringVar(&reportFlagValues.reportTypeFlag, "type", "list", "type od the report")
	set.StringVar(&reportFlagValues.reportTypeFlag, "t", "list", "type od the report")
	set.BoolVar(&reportFlagValues.listTypesFlag, "list-types", false, "list-types")
	set.BoolVar(&reportFlagValues.listTypesFlag, "ls", false, "list-types")
	set.BoolVar(&reportFlagValues.helpFlag, "help", false, "help")
	set.BoolVar(&reportFlagValues.helpFlag, "h", false, "help")

	set.Usage = func() {
		fmt.Print(reportUsageLong)
	}

	return reportCmd{
		repo: repo,
		Cmd:  cmd,
		fs:   set,
	}
}

func (cmd reportCmd) listReport() {
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

func (cmd reportCmd) listReportType() {
	fmt.Print("List report types...")
}

func (cmd reportCmd) Run(args []string) {
	if err := cmd.fs.Parse(args); err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	if *&reportFlagValues.helpFlag {
		cmd.Help()
		return
	}

	if *&reportFlagValues.listTypesFlag {
		cmd.listReportType()
		return
	}

	switch *&reportFlagValues.reportTypeFlag {
	case "list":
		cmd.listReport()
	}
}

func (cmd reportCmd) Help() {
	cmd.fs.Usage()
}

func (cmd reportCmd) ShortHelp() {
	fmt.Print(reportUsageShort)
}
