package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
	repo "totodo/pkg/repository"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
	set.StringVar(&reportFlagValues.reportTypeFlag, "type", "", "type od the report")
	set.StringVar(&reportFlagValues.reportTypeFlag, "t", "", "type od the report")
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
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		HeaderStyle = re.NewStyle().
				Foreground(lipgloss.Color("99")).
				Bold(true).
				Padding(0, 1).
				Align(lipgloss.Center)
		CellStyle = re.NewStyle().
				Padding(0, 1)
	)

	t := table.New().
		Border(lipgloss.NormalBorder()).
		Headers("Task ID", "Description", "Created").
		StyleFunc(func(row, col int) lipgloss.Style {
			var style lipgloss.Style

			switch {
			case row == table.HeaderRow:
				return HeaderStyle
			default:
				style = CellStyle
			}

			return style
		})

	tasks, _ := cmd.repo.GetTasks()

	for _, task := range tasks {
		t.Row(strconv.Itoa(task.Id), task.Description, task.Created.Format(time.DateTime))
	}

	fmt.Println(t)
}

func (cmd reportCmd) listReportTypes() {
	types := `Available list types:
    list      <description>
    timeline  <description>
  `
	fmt.Print(types)
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
		cmd.listReportTypes()
		return
	}

	switch *&reportFlagValues.reportTypeFlag {
	case "list":
		cmd.listReport()
	default:
		cmd.listReportTypes()
	}
}

func (cmd reportCmd) Help() {
	cmd.fs.Usage()
}

func (cmd reportCmd) ShortHelp() {
	fmt.Print(reportUsageShort)
}
