package repository

import (
	"database/sql"
	"totodo/pkg/model"
)

type TrackingRepository interface {
	GetTrackingForTask(taskId int) ([]model.Tracking, error)
	CreateTrackingForTask(taskId int, tracking model.Tracking) error
	UpdateTrackingForTask(taskId int, tracking model.Tracking) error
	DeleteTracking(trackingId int) error
}

type trackingRepository struct {
	db *sql.DB
}

func NewTrackingRepository(db *sql.DB) TrackingRepository {
	return &trackingRepository{db}
}

func (r *trackingRepository) GetTrackingForTask(taskId int) ([]model.Tracking, error) {
	var trackingList []model.Tracking
	return trackingList, nil
}

func (r *trackingRepository) CreateTrackingForTask(taskId int, tracking model.Tracking) error {
	return nil
}

func (r *trackingRepository) UpdateTrackingForTask(taskId int, tracking model.Tracking) error {
	return nil
}

func (r *trackingRepository) DeleteTracking(trackingId int) error {
	return nil
}
