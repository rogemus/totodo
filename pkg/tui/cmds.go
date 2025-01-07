package tui

import (
	"fmt"
	"totodo/pkg/model"

	tea "github.com/charmbracelet/bubbletea"
)

type ChangeToTasksListViewMsg string

func NewChangeToTaskListViewCmd(project model.Project) tea.Cmd {
	projectId := fmt.Sprintf("%d", project.Id)
	return func() tea.Msg {
		return ChangeToTasksListViewMsg(projectId)
	}
}
