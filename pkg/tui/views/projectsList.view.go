package views

import (
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/ui"
	"totodo/pkg/utils"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type projectsListViewModel struct {
	list    list.Model
	project model.Project
	repo    repository.ProjectsRepository
}

func NewProjectsListViewModel(repo repository.ProjectsRepository) projectsListViewModel {
	projects, _ := repo.GetProjects()
	items := utils.ConvertToListitem(projects)

	m := projectsListViewModel{
		list: list.New(items, list.NewDefaultDelegate(), 0, 0),
		repo: repo,
	}

	return m
}

func (m projectsListViewModel) Init() tea.Cmd { return nil }

func (m projectsListViewModel) View() string {
	return ui.WrapperStyle.Render(m.list.View())
}

func (m projectsListViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := ui.WrapperStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)

	case tui.ChangeViewMsg:
		projects, _ := m.repo.GetProjects()
		items := utils.ConvertToListitem(projects)

		m.list.SetItems(items)

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {

		case "esc":
			return m, nil

		case "enter":
			if project, ok := m.list.SelectedItem().(model.Project); ok {
				return m, tea.Batch(tui.NewChangeViewWithProject(project, tui.TASKS_LIST_VIEW), tea.WindowSize())
			} else {
				return m, nil
			}

		case "e":
			project, ok := m.list.SelectedItem().(model.Project)

			if ok {
				return m, tea.Batch(tui.NewChangeViewWithProject(project, tui.CREATE_PROJECT_VIEW), tea.WindowSize())
			}

		case "a":
			return m, tea.Batch(tui.NewChangeViewCmd(tui.CREATE_PROJECT_VIEW), tea.WindowSize())

		case "ctrl+c":
			return m, tea.Quit

		case "X":
			project, ok := m.list.SelectedItem().(model.Project)

			if ok {
				return m, tea.Batch(tui.NewChangeViewWithProject(project, tui.DELETE_PROJECT_VIEW), tea.WindowSize())
			}
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
