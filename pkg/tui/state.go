package tui

import (
	"totodo/pkg/model"
)

type TuiView int

const (
	TASKS_LIST_VIEW TuiView = iota
	DELETE_TASK_VIEW
	CREATE_TASK_VIEW

	PROJECTS_LIST_VIEW
	DELETE_PROJECT_VIEW
	CREATE_PROJECT_VIEW
)

type tuiState struct {
	SelectedTask    model.Task
	SelectedProject model.Project
}

func NewTuiState() tuiState {
	return tuiState{}
}

func (s *tuiState) SetTask(task model.Task) {
	s.SelectedTask = task
}

func (s *tuiState) SetProject(project model.Project) {
	s.SelectedProject = project
}

var State = NewTuiState()
