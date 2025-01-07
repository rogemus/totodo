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
	Description string
	Created     time.Time
	Status      string
	ProjectId      int
	ProjectName    string
}

func NewTask(desc string, listId int) Task {
	created := time.Now()

	return Task{
		Description: desc,
		Created:     created,
		Status:      Status.TODO,
		ProjectId:      listId,
	}
}

func (t *Task) GetListEntry() string {
	id := fmt.Sprintf("[dim]%d[-]", t.Id)
	description := fmt.Sprintf("[cyan]%s[-]", t.Description)
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
