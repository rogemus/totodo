package views

import (
	tea "github.com/charmbracelet/bubbletea"
	"totodo/pkg/repository"
)

type projectsListViewModel struct{}

func NewProjectsListViewModel(listRepo repository.ProjectsRepository) projectsListViewModel {
	return projectsListViewModel{}
}

func (m projectsListViewModel) Init() tea.Cmd {
	return nil
}

func (m projectsListViewModel) View() string {
	return "projectsListViewModel"
}

func (m projectsListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	return m, cmd
}

