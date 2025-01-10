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
)

type tasksListViewModel struct {
	list    list.Model
	repo    repository.TasksRepository
	project model.Project
}

func NewTasksListViewModel(repo repository.TasksRepository) tasksListViewModel {
	m := tasksListViewModel{
		repo:    repo,
		list:    list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		project: model.Project{},
	}

	return m
}

func (m tasksListViewModel) Init() tea.Cmd { return nil }

func (m tasksListViewModel) View() string {
	return ui.WrapperStyle.Render(m.list.View())
}

func (m tasksListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := ui.WrapperStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

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
