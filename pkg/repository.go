package pkg

import (
	"database/sql"
	"fmt"
)

type TasksRepository interface {
	GetTask(id int) (Task, error)
	GetTasks() ([]Task, error)
	UpdateTask(task Task) error
	DeleteTasks(ids []int) error
	CreateTask(task Task) error
}

type tasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) TasksRepository {
	return &tasksRepository{db}
}

func (r *tasksRepository) GetTask(id int) (Task, error) {
	var task Task
	return task, nil
}

func (r *tasksRepository) GetTasks() ([]Task, error) {
	var tasks = make([]Task, 0)
	return tasks, nil
}

func (r *tasksRepository) UpdateTask(task Task) error {
	return nil
}

func (r *tasksRepository) DeleteTasks(ids []int) error {
	return nil
}

func (r *tasksRepository) CreateTask(task Task) error {
	fmt.Printf("added task: %s, +%s, @%s", task.Description, task.Tag, task.Project)
	return nil
}
