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

func (p Project) Title() string {
	return p.Name
}

func (p Project) Description() string {
	return ""
}

func (p Project) FilterValue() string {
	return p.Name
}
