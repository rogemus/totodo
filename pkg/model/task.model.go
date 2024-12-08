package model

import "time"

type Task struct {
	Id          int
	Description string
	Created     time.Time
}

func NewTask(desc string) Task {
	return Task{
		Description: desc,
	}
}
