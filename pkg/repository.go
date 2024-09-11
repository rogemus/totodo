package pkg

import (
	"database/sql"
	"errors"
)

type TasksRepository interface {
	GetTask(id int) (Task, error)
	GetTasks() ([]Task, error)
	UpdateTask(task Task) (int64, error)
	DeleteTask(id int) error
	CreateTask(task Task) (int64, error)
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

func (r *tasksRepository) UpdateTask(task Task) (int64, error) {
	return 0, nil
}

func (r *tasksRepository) DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1;"

	_, err := r.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) CreateTask(task Task) (int64, error) {
	query := "INSERT INTO tasks (description) VALUES ($1)"

	if task.Description == "" {
		return -1, errors.New("empty description")
	}

	result, err := r.db.Exec(query, task.Description)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}
