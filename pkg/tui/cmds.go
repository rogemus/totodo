package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type ChangeViewMsg int

func NewChangeViewCmd(view TuiView) tea.Cmd {
	return func() tea.Msg {
		return ChangeViewMsg(view)
	}
}
