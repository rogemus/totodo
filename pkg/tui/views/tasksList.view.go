package views

import (
	"fmt"
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/ui"
	"totodo/pkg/utils"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tasksListViewModel struct {
	list    list.Model
	repo    repository.TasksRepository
	project model.Project
}

func NewTasksListViewModel(repo repository.TasksRepository) tasksListViewModel {
	list := list.New([]list.Item{}, model.NewTaskItemDelegate(), 0, 0)
	list.SetShowHelp(false)
	list.SetShowStatusBar(false)
	list.SetShowTitle(false)

	m := tasksListViewModel{
		repo:    repo,
		list:    list,
		project: model.Project{},
	}

	return m
}

func (m tasksListViewModel) Init() tea.Cmd { return nil }

func (m tasksListViewModel) View() string {
	listName := fmt.Sprintf("@%s", m.project.Name)
	listMeta := fmt.Sprintf("[%d/%d]", m.project.TasksDoneCount, m.project.TasksCount)
	listTitle := fmt.Sprintf("%s %s", ui.ListNameStyle.Render(listName), ui.DimTextStyle.Render(listMeta))

	tasksDone := fmt.Sprintf("%s %s",
		ui.GreenTextStyle.Render(fmt.Sprintf("%d", m.project.TasksDoneCount)),
		ui.DimTextStyle.Render("done"),
	)
	tasksNotDone := fmt.Sprintf("%s %s",
		ui.MagentaTextStyle.Render(fmt.Sprintf("%d", m.project.TasksCount-m.project.TasksDoneCount)),
		ui.DimTextStyle.Render("pending"),
	)

	// TODO calculate %
	listMetaStats := ui.DimTextStyle.Render("% of the tasks completed.")
	listProgress := lipgloss.JoinHorizontal(
		lipgloss.Top,
		tasksDone,
		ui.DimTextStyle.Render(" | "),
		tasksNotDone,
	)

	return ui.WrapperStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			listTitle,
			m.list.View(),
			lipgloss.JoinVertical(
				lipgloss.Top,
				listMetaStats,
				listProgress,
			),
		),
	)
}

func (m tasksListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := ui.WrapperStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h-3, msg.Height-v-3)

	case tui.ChangeViewWithProjectMsg:
		m.project = msg.Project
		tasks, _ := m.repo.GetTasks(m.project.Id)
		items := utils.ConvertToListitem(tasks)

		m.list.Title = fmt.Sprintf("@%s", m.project.Name)
		m.list.SetItems(items)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "a":
			return m, tea.Batch(tui.NewChangeViewWithProject(m.project, tui.CREATE_TASK_VIEW), tea.WindowSize())

		case "X":
			task, ok := m.list.SelectedItem().(model.Task)

			if ok {
				return m, tea.Batch(tui.NewChangeViewWithTask(task, tui.DELETE_TASK_VIEW), tea.WindowSize())
			}

		case "s":
			return m, tea.Batch(tui.NewChangeViewCmd(tui.PROJECTS_LIST_VIEW), tea.WindowSize())

		case "ctrl+c":
			return m, tea.Quit

		case "esc":
			return m, nil
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
