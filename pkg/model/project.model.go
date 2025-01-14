package model

import (
	"fmt"
	"math"
	"time"
)

type Project struct {
	Id             int
	Name           string
	TasksCount     int
	TasksDoneCount int
	Created        time.Time
}

func NewProject(name string) Project {
	created := time.Now()

	return Project{
		Name:    name,
		Created: created,
	}
}

func (p *Project) Stat() int {
	if p.TasksCount == 0 {
		return 0
	}

	stat := (p.TasksDoneCount / p.TasksCount) * 100
	statFloat := math.Round(float64(stat))
	return int(statFloat)
}

func (p *Project) GetTimeSinceCreation() string {
	since := time.Since(p.Created)
	hours := since.Hours()
	days := int(hours / 24)

	if days < 1 {
		return fmt.Sprintf("%dh", int(hours))
	}

	return fmt.Sprintf("%dd", days)
}

func (p Project) FilterValue() string {
	return p.Name
}
