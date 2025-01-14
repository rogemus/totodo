package ui

import "github.com/charmbracelet/lipgloss"

var DimTextStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Dim)

var CyanTextStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Cyan)

var GreenTextStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Green)

var RedTextStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Red)

var BlueTextStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Blue)

var MagentaTextStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Magenta)

var YellowTextStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Yellow)

var EntryTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(NormalColors.Magenta)

var EntryStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	MarginLeft(0)

var ListNameStyle = lipgloss.NewStyle().
	Foreground(NormalColors.Blue).
	Underline(true)

var ListTitleStyle = lipgloss.NewStyle().
	PaddingLeft(1).
	MarginLeft(0).
	MarginBottom(2)

var BlankStyle = lipgloss.NewStyle()
