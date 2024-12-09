package cmd

import (
	"flag"
	"fmt"
	"time"
	repo "totodo/pkg/repository"
)

type editCmd struct {
	Cmd  string
	repo repo.TasksRepository
	fs   *flag.FlagSet
}

var editFlagValues = struct {
	idFlag   int
	descFlag string
	helpFlag bool
}{
	idFlag:   0,
	descFlag: "",
	helpFlag: false,
}

const editUsageLong = `
  Edit task:
    -i, --id    id of the task
    -d, --desc  description of the task
    -h, --help  prints help information
`
const editUsageShort = `
  Edit task:
    -i, --id    id of the task
    -d, --desc  description of the task
`

func NewEditCmd(repo repo.TasksRepository) editCmd {
	cmd := "edit"
	set := flag.NewFlagSet(cmd, flag.ContinueOnError)

	set.IntVar(&editFlagValues.idFlag, "id", 0, "`id` of the task")
	set.IntVar(&editFlagValues.idFlag, "i", 0, "`id` of the task")
	set.BoolVar(&editFlagValues.helpFlag, "help", false, "help")
	set.BoolVar(&editFlagValues.helpFlag, "h", false, "help")

	set.Usage = func() {
		fmt.Print(editUsageLong)
	}

	return editCmd{
		repo: repo,
		Cmd:  cmd,
		fs:   set,
	}
}

func (cmd editCmd) Run(args []string) {
	if err := cmd.fs.Parse(args); err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	if *&editFlagValues.helpFlag {
		cmd.Help()
		return
	}

	if *&editFlagValues.idFlag == 0 {
		fmt.Print("Id not provided. Available options: \n\n")
		cmd.Help()
		return
	}

	if *&editFlagValues.descFlag == "" {
		fmt.Print("No description provided. Available options: \n\n")
		cmd.Help()
		return
	}

	task, err := cmd.repo.GetTask(*&editFlagValues.idFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	task.Description = *&editFlagValues.descFlag
	err = cmd.repo.UpdateTask(task)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Task Updated: [%d] (@%s) %s", task.Id, task.Created.Format(time.DateTime), task.Description)
}

func (cmd editCmd) Help() {
	cmd.fs.Usage()
}

func (cmd editCmd) ShortHelp() {
	fmt.Print(editUsageShort)
}
