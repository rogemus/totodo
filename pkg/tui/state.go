package tui

type TuiView int

const (
	TASKS_LIST_VIEW TuiView = iota
	DELETE_TASK_VIEW
	CREATE_TASK_VIEW

	PROJECTS_LIST_VIEW
	DELETE_PROJECT_VIEW
	CREATE_PROJECT_VIEW
)
