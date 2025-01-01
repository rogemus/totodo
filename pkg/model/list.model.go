package model

import (
	"time"
)

type List struct {
	Id      int
	Name    string
	Created time.Time
}

func NewList(name string) List {
	created := time.Now()

	return List{
		Name:    name,
		Created: created,
	}
}
