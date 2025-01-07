package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"totodo/pkg/repository"
)

type createProjectViewModel struct{}

func NewCreateProjectViewModel(listRepo repository.ProjectsRepository) createProjectViewModel {
	return createProjectViewModel{}
}

func (m createProjectViewModel) Init() tea.Cmd { return nil }

func (m createProjectViewModel) View() string {
	return "createProjectViewModel"
}

func (m createProjectViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}
