package cmd

import (
	"flag"
	"fmt"
	repo "totodo/pkg/repository"
)

type deleteCmd struct {
	Cmd  string
	repo repo.TasksRepository
	fs   *flag.FlagSet
}

var deleteFlagValues = struct {
	idFlag   int
	helpFlag bool
}{
	idFlag:   0,
	helpFlag: false,
}

const deleteUsageLong = `
  Delete task:
    -i, --id    id of the task
    -h, --help  prints help information
`
const deleteUsageShort = `
  Delete task:
    -i, --id    id of the task
`

func NewDeleteCmd(repo repo.TasksRepository) deleteCmd {
	cmd := "delete"
	set := flag.NewFlagSet(cmd, flag.ContinueOnError)

	set.IntVar(&deleteFlagValues.idFlag, "id", 0, "`id` of the task")
	set.IntVar(&deleteFlagValues.idFlag, "i", 0, "`id` of the task")
	set.BoolVar(&deleteFlagValues.helpFlag, "help", false, "help")
	set.BoolVar(&deleteFlagValues.helpFlag, "h", false, "help")

	set.Usage = func() {
		fmt.Print(deleteUsageLong)
	}

	return deleteCmd{
		repo: repo,
		Cmd:  cmd,
		fs:   set,
	}
}

func (cmd deleteCmd) Run(args []string) {
	if err := cmd.fs.Parse(args); err != nil {
		fmt.Printf("error: %s", err)
		return
	}

	if *&deleteFlagValues.helpFlag {
		cmd.Help()
		return
	}

	if *&deleteFlagValues.idFlag == 0 {
		fmt.Print("Id not provided. Available options: \n\n")
		cmd.Help()
		return
	}

	task, err := cmd.repo.GetTask(*&deleteFlagValues.idFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = cmd.repo.DeleteTask(*&deleteFlagValues.idFlag)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Task Deleted: [%d]", task.Id)
}

func (cmd deleteCmd) Help() {
	cmd.fs.Usage()
}

func (cmd deleteCmd) ShortHelp() {
	fmt.Print(deleteUsageShort)
}
