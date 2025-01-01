package cmd

import (
	"flag"
	"fmt"
	"time"
	"totodo/pkg/model"
	repo "totodo/pkg/repository"
)

type addCmd struct {
	Cmd  string
	repo repo.TasksRepository
	fs   *flag.FlagSet
}

var addFlagValues = struct {
	descFlag string
	helpFlag bool
}{
	descFlag: "",
	helpFlag: false,
}

const addUsageLong = `
  Create task:
    -d, --desc  description of the task
    -h, --help  prints help information
`
const addUsageShort = `
  Create task:
    -d, --desc  description of the task
`

func NewAddCmd(repo repo.TasksRepository) addCmd {
	cmd := "add"
	set := flag.NewFlagSet(cmd, flag.ContinueOnError)

	set.StringVar(&addFlagValues.descFlag, "desc", "", "`description` of the task")
	set.StringVar(&addFlagValues.descFlag, "d", "", "`description` of the task")
	set.BoolVar(&addFlagValues.helpFlag, "help", false, "help")
	set.BoolVar(&addFlagValues.helpFlag, "h", false, "help")

	set.Usage = func() {
		fmt.Print(addUsageLong)
	}

	return addCmd{
		repo: repo,
		Cmd:  cmd,
		fs:   set,
	}
}

func (cmd addCmd) Run(args []string) {
	if err := cmd.fs.Parse(args); err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	if *&addFlagValues.helpFlag {
		cmd.Help()
		return
	}

	if *&addFlagValues.descFlag == "" {
		fmt.Print("No description provided. Available options: \n\n")
		cmd.Help()
		return
	}

  // TODO: addy dynaminc listId
	task := model.NewTask(*&addFlagValues.descFlag, 1)
	taskId, err := cmd.repo.CreateTask(task)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Task Created: [%d] (@%s) %s", taskId, task.Created.Format(time.DateTime), task.Description)
}

func (cmd addCmd) Help() {
	cmd.fs.Usage()
}

func (cmd addCmd) ShortHelp() {
	fmt.Print(addUsageShort)
}
