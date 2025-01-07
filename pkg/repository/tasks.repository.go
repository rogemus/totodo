package repository

import (
	"database/sql"
	_ "embed"
	"errors"
	"totodo/pkg/model"
)

var (
	//go:embed queries/tasks/getTask.sql
	getTaskQuery string
	//go:embed queries/tasks/getTasks.sql
	getTasksQuery string
	//go:embed queries/tasks/updateTask.sql
	updateTaskQuery string
	//go:embed queries/tasks/deleteTask.sql
	deleteTaskQuery string
	//go:embed queries/tasks/createTask.sql
	createTaskQuery string
)

type TasksRepository interface {
	GetTask(id int) (model.Task, error)
	GetTasks() ([]model.Task, error)
	UpdateTask(task model.Task) error
	DeleteTask(id int) error
	CreateTask(task model.Task) (int64, error)
}

type tasksRepository struct {
	db *sql.DB
}

func NewTasksRepository(db *sql.DB) TasksRepository {
	return &tasksRepository{db}
}

func (r *tasksRepository) GetTask(id int) (model.Task, error) {
	var task model.Task

	row := r.db.QueryRow(getTaskQuery, id)
	err := row.Scan(
		&task.Id,
		&task.Description,
		&task.Created,
		&task.Status,
		&task.ProjectId,
		&task.ProjectName,
	)

	if err != nil {
		return task, err
	}

	return task, nil
}

func (r *tasksRepository) GetTasks() ([]model.Task, error) {
	rows, err := r.db.Query(getTasksQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tasks := make([]model.Task, 0)

	for rows.Next() {
		var task model.Task

		err := rows.Scan(
			&task.Id,
			&task.Description,
			&task.Created,
			&task.Status,
			&task.ProjectId,
			&task.ProjectName,
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

func (r *tasksRepository) UpdateTask(task model.Task) error {
	_, err := r.db.Exec(updateTaskQuery, task.Description, task.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) DeleteTask(id int) error {
	_, err := r.db.Exec(deleteTaskQuery, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *tasksRepository) CreateTask(task model.Task) (int64, error) {
	if task.Description == "" {
		return -1, errors.New("empty description")
	}

	result, err := r.db.Exec(createTaskQuery, task.Description, task.Status, task.ProjectId)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}
