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
}

func NewTask(desc string) Task {
	created := time.Now()

	return Task{
		Description: desc,
		Created:     created,
		Status:      Status.TODO,
	}
}

func (t *Task) GetStatusIcon() string {
	// if t.Status == Status.ACTIVE {
	// 	return "★"
	// }

	if t.Status == Status.DONE {
		return "✓"
	}

	return "☐"
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
