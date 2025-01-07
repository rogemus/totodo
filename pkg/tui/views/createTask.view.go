package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"totodo/pkg/repository"
)

type createTaskViewModel struct{}

func NewCreateTaskViewModel(repo repository.TasksRepository) createTaskViewModel {
	return createTaskViewModel{}
}

func (m createTaskViewModel) Init() tea.Cmd {
	return nil
}

func (m createTaskViewModel) View() string {
	return "createTaskViewModel"
}

func (m createTaskViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

