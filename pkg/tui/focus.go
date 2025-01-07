package tui

type Focus int

const (
	CANCEL_BTN Focus = iota
	CONFIRM_BTN

	NAME_INPUT
	DESCRIPTION_INPUT
	CREATED_INPUT
)
