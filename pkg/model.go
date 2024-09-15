package pkg

import "time"

type Cmd interface {
	Run(args []string)
	Help()
}

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
