package ui

import "github.com/charmbracelet/lipgloss"

var DialogTitleStyle = lipgloss.NewStyle().
	Bold(true).
	MarginTop(2)

var DialogBoxStyle = lipgloss.NewStyle().
	Padding(1).
	Border(lipgloss.RoundedBorder())

var DialogTextStyle = lipgloss.NewStyle()

var DialogFooterStyle = lipgloss.NewStyle().
	MarginTop(1)

var ButtonStyle = lipgloss.NewStyle().
	Padding(0, 3)

var ConfirmBtnStyle = ButtonStyle.
	Foreground(NormalColors.White).
	// Background(NormalColors.Green).
	MarginLeft(2)

var CancelBtnStyle = ButtonStyle.
	Foreground(NormalColors.White)
	// Background(NormalColors.Red)
