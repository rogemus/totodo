package cmd

import (
	"fmt"
	"os"
	"totodo/pkg/repository"
	"totodo/pkg/tui/views"

	tea "github.com/charmbracelet/bubbletea"
)

type selectedView int

const (
	TASKS_LIST_VIEW selectedView = iota
	DELETE_TASK_VIEW
	CREATE_TASK_VIEW

	PROJECTS_LIST_VIEW
	DELETE_PROJECT_VIEW
	CREATE_PROJECT_VIEW
)

type TUIModel struct {
	projectsListModel  tea.Model
	createProjectModel tea.Model
	deleteProjectModel tea.Model
	tasksListModel     tea.Model
	createTaskModel    tea.Model
	deleteTaskModel    tea.Model

	selectedView selectedView
}

func NewTui(projectsRepo repository.ProjectsRepository, tasksRepo repository.TasksRepository) TUIModel {
	return TUIModel{
		projectsListModel:  views.NewProjectsListViewModel(projectsRepo),
		createProjectModel: views.NewCreateProjectViewModel(projectsRepo),
		deleteProjectModel: views.NewDeleteProjectViewModel(projectsRepo),
		tasksListModel:     views.NewTasksListViewModel(projectsRepo),
		createTaskModel:    views.NewCreateTaskViewModel(projectsRepo),
		deleteTaskModel:    views.NewDeleteTaskViewModel(projectsRepo),

		selectedView: PROJECTS_LIST_VIEW,
	}
}

func (m TUIModel) Init() tea.Cmd {
	return nil
}

func (m TUIModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			m.selectedView = CREATE_TASK_VIEW

		case "A":
			m.selectedView = CREATE_PROJECT_VIEW

		case "s":
			m.selectedView = PROJECTS_LIST_VIEW

		case "t":
			m.selectedView = TASKS_LIST_VIEW

		case "x":
			m.selectedView = DELETE_TASK_VIEW

		case "X":
			m.selectedView = DELETE_PROJECT_VIEW
		}
	}

	switch m.selectedView {
	case PROJECTS_LIST_VIEW:
		m.projectsListModel, cmd = m.projectsListModel.Update(msg)

	case CREATE_PROJECT_VIEW:
		m.createProjectModel, cmd = m.createProjectModel.Update(msg)

	case DELETE_PROJECT_VIEW:
		m.deleteProjectModel, cmd = m.deleteProjectModel.Update(msg)

	case TASKS_LIST_VIEW:
		m.tasksListModel, cmd = m.tasksListModel.Update(msg)

	case CREATE_TASK_VIEW:
		m.createTaskModel, cmd = m.createTaskModel.Update(msg)

	case DELETE_TASK_VIEW:
		m.deleteTaskModel, cmd = m.deleteTaskModel.Update(msg)
	}

	return m, cmd
}

func (m TUIModel) View() string {
	switch m.selectedView {
	case PROJECTS_LIST_VIEW:
		return m.projectsListModel.View()

	case CREATE_PROJECT_VIEW:
		return m.createProjectModel.View()

	case DELETE_PROJECT_VIEW:
		return m.deleteProjectModel.View()

	case TASKS_LIST_VIEW:
		return m.tasksListModel.View()

	case CREATE_TASK_VIEW:
		return m.createTaskModel.View()

	case DELETE_TASK_VIEW:
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
