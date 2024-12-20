package ui

import (
	"github.com/charmbracelet/lipgloss"
)

type colors struct {
	Black   lipgloss.Color
	Red     lipgloss.Color
	Green   lipgloss.Color
	Yellow  lipgloss.Color
	Blue    lipgloss.Color
	Magenta lipgloss.Color
	Cyan    lipgloss.Color
	White   lipgloss.Color
	Dim     lipgloss.Color
}

var NormalColors = colors{
	Black:   lipgloss.Color("0"),
	Red:     lipgloss.Color("1"),
	Green:   lipgloss.Color("2"),
	Yellow:  lipgloss.Color("3"),
	Blue:    lipgloss.Color("4"),
	Magenta: lipgloss.Color("5"),
	Cyan:    lipgloss.Color("6"),
	White:   lipgloss.Color("7"),
	Dim:     lipgloss.Color("240"),
}

var BrightColors = colors{
	Black:   lipgloss.Color("8"),
	Red:     lipgloss.Color("9"),
	Green:   lipgloss.Color("10"),
	Yellow:  lipgloss.Color("11"),
	Blue:    lipgloss.Color("12"),
	Magenta: lipgloss.Color("13"),
	Cyan:    lipgloss.Color("14"),
	White:   lipgloss.Color("15"),
}
