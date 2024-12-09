package model

import "time"

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
