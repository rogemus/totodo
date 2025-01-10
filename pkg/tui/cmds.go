package tui

import (
	"log"
	"totodo/pkg/model"

	tea "github.com/charmbracelet/bubbletea"
)

type ChangeViewMsg struct {
	View TuiView
}

type ChangeViewWithTaskMsg struct {
	ChangeViewMsg
	Task model.Task
}

type ChangeViewWithProjectMsg struct {
	ChangeViewMsg
	Project model.Project
}

func NewChangeViewWithTask(task model.Task, view TuiView) tea.Cmd {
	log.Printf("[INFO] View change: Task - [%d] %s | View - %d", task.Id, task.Name, view)

	return func() tea.Msg {
		return ChangeViewWithTaskMsg{Task: task, ChangeViewMsg: ChangeViewMsg{View: view}}
	}
}

func NewChangeViewWithProject(project model.Project, view TuiView) tea.Cmd {
	log.Printf("[INFO] View change: Project - [%d] %s | View - %d", project.Id, project.Name, view)

	return func() tea.Msg {
		return ChangeViewWithProjectMsg{Project: project, ChangeViewMsg: ChangeViewMsg{View: view}}
	}
}

func NewChangeViewCmd(view TuiView) tea.Cmd {
	log.Printf("[INFO] View change: View - %d", view)

	return func() tea.Msg {
		return ChangeViewMsg{View: view}
	}
}
