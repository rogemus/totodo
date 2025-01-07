package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"totodo/pkg/repository"
)

type deleteProjectViewModel struct{}

func NewDeleteProjectViewModel(listRepo repository.ProjectsRepository) deleteProjectViewModel {
	return deleteProjectViewModel{}
}

func (m deleteProjectViewModel) Init() tea.Cmd { return nil }

func (m deleteProjectViewModel) View() string {
	return "deleteProjectViewModel"
}

func (m deleteProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}
