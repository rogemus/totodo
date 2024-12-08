package model

import "time"

type Tracking struct {
	Id          int
	Created     time.Time
}

func NewTracking(desc string) Tracking {
	return Tracking{}
}
