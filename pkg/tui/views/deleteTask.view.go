package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"totodo/pkg/repository"
)

type deleteTaskViewModel struct{}

func NewDeleteTaskViewModel(listRepo repository.ProjectsRepository) deleteTaskViewModel {
	return deleteTaskViewModel{}
}

func (m deleteTaskViewModel) Init() tea.Cmd {
	return nil
}

func (m deleteTaskViewModel) View() string {
	return "deleteTaskViewModel"
}

func (m deleteTaskViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

