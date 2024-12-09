package cmd

import (
	"flag"
	"fmt"
	"os"
	"text/template"
	repo "totodo/pkg/repository"
)

type showCmd struct {
	Cmd  string
	repo repo.TasksRepository
	fs   flag.FlagSet
}

var showFlagValues = struct {
	idFlag   int
	helpFlag bool
}{
	idFlag:   0,
	helpFlag: false,
}

const showUsageLong = `
  Show task:
    -i, --id    id of the task
    -h, --help  prints  help information
`
const showUsageShort = `
  Show task:
    -i, --id    id of the task
`

func NewShowCmd(repo repo.TasksRepository) showCmd {
	cmd := "show"
	set := flag.NewFlagSet(cmd, flag.ContinueOnError)

	set.IntVar(&showFlagValues.idFlag, "id", 0, "`id` of the task")
	set.IntVar(&showFlagValues.idFlag, "i", 0, "`id` of the task")
	set.BoolVar(&showFlagValues.helpFlag, "help", false, "help")
	set.BoolVar(&showFlagValues.helpFlag, "h", false, "help")

	set.Usage = func() {
		fmt.Print(showUsageLong)
	}

	return showCmd{
		repo: repo,
		Cmd:  "show",
		fs:   *set,
	}
}

func (cmd showCmd) showTask() {
	// TODO use lipgoss for proper template
	tmpl :=
		`
      id  |  description  |  created
    ----------------------------------
      {{ .Id }}  |  {{ .Description }}  |  {{ .Created }}
    `
	task, err := cmd.repo.GetTask(showFlagValues.idFlag)

	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	t := template.Must(template.New("list").Parse(tmpl))

	if err := t.Execute(os.Stdout, task); err != nil {
		fmt.Printf("%v", err)
	}
}

func (cmd showCmd) Run(args []string) {
	if err := cmd.fs.Parse(args); err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	if *&showFlagValues.helpFlag {
		cmd.Help()
		return
	}

	if *&showFlagValues.idFlag == 0 {
		fmt.Print("Id not provided. Available options: \n\n")
		cmd.Help()
		return
	}

	cmd.showTask()
}

func (cmd showCmd) Help() {
	cmd.fs.Usage()
}

func (cmd showCmd) ShortHelp() {
	fmt.Print(showUsageShort)
}
