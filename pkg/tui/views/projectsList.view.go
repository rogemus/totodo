package views

import (
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/utils"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type projectsListViewModel struct {
	list    list.Model
	project model.Project
}

func NewProjectsListViewModel(repo repository.ProjectsRepository) projectsListViewModel {
	projects, _ := repo.GetProjects()
	items := utils.ConvertToListitem(projects)

	m := projectsListViewModel{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}

	return m
}

func (m projectsListViewModel) Init() tea.Cmd { return nil }

func (m projectsListViewModel) View() string {
	return m.list.View()
}

func (m projectsListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "enter":
			project, ok := m.list.SelectedItem().(model.Project)

			if ok {
				tui.State.SetProject(project)
			}

			cmd = tui.NewChangeToTaskListViewCmd(project)
			return m, tea.Batch(cmd, tea.WindowSize())
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
