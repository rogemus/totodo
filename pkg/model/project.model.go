package model

import (
	"time"
)

type Project struct {
	Id      int
	Name    string
	Created time.Time
}

func NewProject(name string) Project {
	created := time.Now()

	return Project{
		Name:    name,
		Created: created,
	}
}
