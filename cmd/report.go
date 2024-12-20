package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"totodo/pkg/model"
	repo "totodo/pkg/repository"
	"totodo/pkg/ui"

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

func GroupByStatus(tasks []model.Task) ([]model.Task, []model.Task, []model.Task) {
	doneTasks := make([]model.Task, 0)
	todoTasks := make([]model.Task, 0)
	activeTasks := make([]model.Task, 0)

	for _, task := range tasks {
		switch task.Status {
		case model.Status.ACTIVE:
			activeTasks = append(activeTasks, task)

		case model.Status.TODO:
			todoTasks = append(todoTasks, task)

		case model.Status.DONE:
			doneTasks = append(doneTasks, task)
		}
	}

	return activeTasks, todoTasks, doneTasks
}

func (cmd reportCmd) listReport() {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		CellStyle    = re.NewStyle().Padding(0, 1)
		IndexStyle   = CellStyle.Foreground(ui.NormalColors.Dim).Width(3)
		CreatedStyle = CellStyle.Foreground(ui.NormalColors.Dim)

		TodoStatusStyles = CellStyle.Foreground(ui.BrightColors.Blue).Width(3)
		TodoTitleStyles  = CellStyle.Foreground(ui.NormalColors.Blue)

		DoneStatusStyles = CellStyle.Foreground(ui.NormalColors.Green).Width(3)
		DoneTitleStyles  = CellStyle.Foreground(ui.NormalColors.Dim).Strikethrough(true)

		ActiveStatusStyles = re.NewStyle().Foreground(ui.BrightColors.Yellow).Width(1)
		ActiveTitleStyles  = CellStyle.Foreground(ui.NormalColors.Yellow).Underline(true)

		GreenTextStyles  = re.NewStyle().Foreground(ui.NormalColors.Green)
		YellowTextStyles = re.NewStyle().Foreground(ui.NormalColors.Yellow)
		BlueTextStyles   = re.NewStyle().Foreground(ui.NormalColors.Blue)
		DimTextStyles    = re.NewStyle().Foreground(ui.NormalColors.Dim)
	)

	tasks, _ := cmd.repo.GetTasks()
	activeTasks, todoTasks, doneTasks := GroupByStatus(tasks)

	t := table.New().Border(lipgloss.HiddenBorder())

	for _, task := range activeTasks {
		idCol := IndexStyle.Render(strconv.Itoa(task.Id))
		statusIcon := task.GetStatusIcon()
		createdCol := CreatedStyle.Render(task.GetTimeSinceCreation())
		statusCol := TodoStatusStyles.Render(statusIcon)
		titleCol := fmt.Sprintf("%s%s", ActiveTitleStyles.Render(task.Description), ActiveStatusStyles.Render("★"))

		t.Row(idCol, statusCol, titleCol, createdCol)
	}

	t.Row("", "", "", "")

	for _, task := range todoTasks {
		idCol := IndexStyle.Render(strconv.Itoa(task.Id))
		statusIcon := task.GetStatusIcon()
		createdCol := CreatedStyle.Render(task.GetTimeSinceCreation())
		statusCol := TodoStatusStyles.Render(statusIcon)
		titleCol := TodoTitleStyles.Render(task.Description)

		t.Row(idCol, statusCol, titleCol, createdCol)
	}

	for _, task := range doneTasks {
		idCol := IndexStyle.Render(strconv.Itoa(task.Id))
		statusIcon := task.GetStatusIcon()
		createdCol := CreatedStyle.Render(task.GetTimeSinceCreation())
		statusCol := DoneStatusStyles.Render(statusIcon)
		titleCol := DoneTitleStyles.Render(task.Description)

		t.Row(idCol, statusCol, titleCol, createdCol)
	}

	separatot := "⋅"
	activeCount := YellowTextStyles.Render(fmt.Sprintf("%d", len(activeTasks)))
	activeLabel := DimTextStyles.Render(fmt.Sprintf("active %s", separatot))
	todoCount := BlueTextStyles.Render(fmt.Sprintf("%d", len(todoTasks)))
	todoLabel := DimTextStyles.Render(fmt.Sprintf("pending %s", separatot))
	doneCount := GreenTextStyles.Render(fmt.Sprintf("%d", len(doneTasks)))
	doneLabel := DimTextStyles.Render(fmt.Sprintf("done %s", separatot))

	fmt.Println(t)
	fmt.Println(CellStyle.Render(fmt.Sprintf(
		"%s %s %s %s %s %s",
		activeCount,
		activeLabel,
		todoCount,
		todoLabel,
		doneCount,
		doneLabel,
	)))
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
