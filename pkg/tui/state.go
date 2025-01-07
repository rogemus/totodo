package tui

import (
	"totodo/pkg/model"
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
