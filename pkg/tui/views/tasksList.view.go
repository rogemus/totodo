package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"totodo/pkg/repository"
)

type tasksListViewModel struct{}

func NewTasksListViewModel(listRepo repository.ProjectsRepository) tasksListViewModel {
	return tasksListViewModel{}
}

func (m tasksListViewModel) Init() tea.Cmd {
	return nil
}

func (m tasksListViewModel) View() string {
	return "tasksListViewModel"
}

func (m tasksListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

