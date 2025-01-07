package views

import (
	"fmt"
	"io"
	"strings"
	"totodo/pkg/model"
	"totodo/pkg/repository"
	"totodo/pkg/tui"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type projectsListViewModel struct {
	list    list.Model
	project model.Project
}

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int {
	return 1
}

func (d itemDelegate) Spacing() int {
	return 0
}

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd {
	return nil
}

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	project, ok := listItem.(model.Project)

	if !ok {
		return
	}

	str := fmt.Sprintf("%d. [%d] %s", index+1, project.Id, project.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return selectedItemStyle.Render("> " + strings.Join(s, " "))
		}
	}

	fmt.Fprint(w, fn(str))
}

func NewProjectsListViewModel(listRepo repository.ProjectsRepository) projectsListViewModel {
	projects := []list.Item{
		model.Project{Name: "Project 1"},
		model.Project{Name: "Project 2"},
	}

	m := projectsListViewModel{
		list: list.New(projects, itemDelegate{}, 0, 0),
	}

	return m
}

func (m projectsListViewModel) Init() tea.Cmd {
	return nil
}

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
			return m, cmd
		}
	}

	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
