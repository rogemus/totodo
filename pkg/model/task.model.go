package model

import (
	"fmt"
	"time"
)

type TaskStatus struct {
	ACTIVE string
	DONE   string
	TODO   string
}

var Status TaskStatus = TaskStatus{
	TODO:   "todo",
	DONE:   "done",
	ACTIVE: "active",
}

type Task struct {
	Id          int
	Name        string
	Created     time.Time
	Status      string
	ProjectId   int
	ProjectName string
}

func NewTask(name string, projectId int) Task {
	created := time.Now()

	return Task{
		Name:      name,
		Created:   created,
		Status:    Status.TODO,
		ProjectId: projectId,
	}
}

func (t Task) FilterValue() string {
	return t.Name
}

func (t Task) Title() string {
	return t.Name
}

func (t Task) Description() string {
	return ""
}

func (t *Task) GetListEntry() string {
	id := fmt.Sprintf("[dim]%d[-]", t.Id)
	description := fmt.Sprintf("[cyan]%s[-]", t.Name)
	created := t.GetEntryCreation()
	entry := fmt.Sprintf("%s %s %s", id, description, created)
	return entry
}

func (t *Task) GetEntryStatus() string {
	if t.Status == Status.DONE {
		return "[green]✓[-]"
	}

	return "[magenta]☐[-]"
}

func (t *Task) GetStatusIcon() string {
	if t.Status == Status.DONE {
		return "✓"
	}

	return "☐"
}

func (t *Task) GetEntryCreation() string {
	time := t.GetTimeSinceCreation()
	return fmt.Sprintf("[dim]%s[-]", time)
}

func (t *Task) GetTimeSinceCreation() string {
	since := time.Since(t.Created)
	hours := since.Hours()
	days := int(hours / 24)

	if days < 1 {
		return fmt.Sprintf("%dh", int(hours))
	}

	return fmt.Sprintf("%dd", days)
}
