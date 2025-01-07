package cmd

import (
	"fmt"
	"os"
	"totodo/pkg/repository"
	"totodo/pkg/tui"
	"totodo/pkg/tui/views"

	tea "github.com/charmbracelet/bubbletea"
)

type TUIModel struct {
	projectsListModel  tea.Model
	createProjectModel tea.Model
	deleteProjectModel tea.Model
	tasksListModel     tea.Model
	createTaskModel    tea.Model
	deleteTaskModel    tea.Model

	selectedView tui.TuiView
}

func NewTui(projectsRepo repository.ProjectsRepository, tasksRepo repository.TasksRepository) TUIModel {
	return TUIModel{
		projectsListModel:  views.NewProjectsListViewModel(projectsRepo),
		createProjectModel: views.NewCreateProjectViewModel(projectsRepo),
		deleteProjectModel: views.NewDeleteProjectViewModel(projectsRepo),
		tasksListModel:     views.NewTasksListViewModel(tasksRepo),
		createTaskModel:    views.NewCreateTaskViewModel(tasksRepo),
		deleteTaskModel:    views.NewDeleteTaskViewModel(tasksRepo),

		selectedView: tui.PROJECTS_LIST_VIEW,
	}
}

func (m TUIModel) Init() tea.Cmd { return nil }

func (m TUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tui.ChangeViewMsg:
		m.selectedView = tui.TuiView(msg)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	switch m.selectedView {
	case tui.PROJECTS_LIST_VIEW:
		m.projectsListModel, cmd = m.projectsListModel.Update(msg)

	case tui.CREATE_PROJECT_VIEW:
		m.createProjectModel, cmd = m.createProjectModel.Update(msg)

	case tui.DELETE_PROJECT_VIEW:
		m.deleteProjectModel, cmd = m.deleteProjectModel.Update(msg)

	case tui.TASKS_LIST_VIEW:
		m.tasksListModel, cmd = m.tasksListModel.Update(msg)

	case tui.CREATE_TASK_VIEW:
		m.createTaskModel, cmd = m.createTaskModel.Update(msg)

	case tui.DELETE_TASK_VIEW:
		m.deleteTaskModel, cmd = m.deleteTaskModel.Update(msg)
	}

	return m, cmd
}

func (m TUIModel) View() string {
	switch m.selectedView {
	case tui.PROJECTS_LIST_VIEW:
		return m.projectsListModel.View()

	case tui.CREATE_PROJECT_VIEW:
		return m.createProjectModel.View()

	case tui.DELETE_PROJECT_VIEW:
		return m.deleteProjectModel.View()

	case tui.TASKS_LIST_VIEW:
		return m.tasksListModel.View()

	case tui.CREATE_TASK_VIEW:
		return m.createTaskModel.View()

	case tui.DELETE_TASK_VIEW:
		return m.deleteTaskModel.View()
	}

	return m.createProjectModel.View()
}

func (m TUIModel) Run() {
	p := tea.NewProgram(&m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Ups ...")
		os.Exit(1)
	}
}
