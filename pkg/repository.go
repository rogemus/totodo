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

	query := "SELECT id, description, created FROM tasks WHERE id = $1;"

	row := r.db.QueryRow(query, id)
	err := row.Scan(
		&task.Id,
		&task.Description,
		&task.Created,
	)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *tasksRepository) GetTasks() ([]Task, error) {
	query := "SELECT id, description, created FROM tasks;"
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tasks := make([]Task, 0)

	for rows.Next() {
		var task Task

		err := rows.Scan(
			&task.Id,
			&task.Description,
			&task.Created,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (r *tasksRepository) UpdateTask(task Task) error {
	return nil
}

func (r *tasksRepository) DeleteTasks(ids []int) error {
	return nil
}

func (r *tasksRepository) CreateTask(task Task) error {
	fmt.Printf("added task: %s, at %v", task.Description, task.Created)
	return nil
}
