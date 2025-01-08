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
	list list.Model
	done bool
	repo repository.TasksRepository
}

func NewTasksListViewModel(repo repository.TasksRepository) tasksListViewModel {
	m := tasksListViewModel{
		repo: repo,
		list: list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
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

	case tui.ChangeViewMsg:
		project := tui.State.SelectedProject
		tasks, _ := m.repo.GetTasks(project.Id)
		items := utils.ConvertToListitem(tasks)

		m.list.Title = fmt.Sprintf("@%s", project.Name)
		m.list.SetItems(items)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "a":
			return m, tea.Batch(tui.NewChangeViewCmd(tui.CREATE_TASK_VIEW), tea.WindowSize())

		case "X":
			task, ok := m.list.SelectedItem().(model.Task)

			if ok {
				tui.State.SetTask(task)
			}

			return m, tea.Batch(tui.NewChangeViewCmd(tui.DELETE_TASK_VIEW), tea.WindowSize())

		case "esc":
			return m, nil
			// do nothing
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
