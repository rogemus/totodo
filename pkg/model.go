package pkg

import "time"

type Task struct {
	Id          int
	Description string
	Created     time.Time
}

func NewTask(desc string) Task {
	created := time.Now()

	return Task{
		Created:     created,
		Description: desc,
	}
}
