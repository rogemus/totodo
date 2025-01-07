package views

import (
	"fmt"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
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
		list: list.New([]list.Item{}, list.NewDefaultDelegate(), 25, 25),
	}

	return m
}

func (m tasksListViewModel) Init() tea.Cmd { return nil }

func (m tasksListViewModel) View() string {
	return docStyle.Render(m.list.View())
}

func (m tasksListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tui.ChangeToTasksListViewMsg:
		project := tui.State.SelectedProject
		tasks, _ := m.repo.GetTasks(project.Id)
		items := utils.ConvertToListitem(tasks)

		m.list.Title = fmt.Sprintf("@%s", project.Name)
		m.list.SetItems(items)

		return m, nil
	}

	return m, cmd
}
