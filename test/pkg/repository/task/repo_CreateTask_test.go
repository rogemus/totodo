package pkg_test

import (
	"errors"
	"testing"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateTask(t *testing.T) {
	var empty_task model.Task

	task := model.Task{
		Name:      "test task",
		ProjectId: 0,
	}

	testCases := []struct {
		name        string
		task        model.Task
		taskId      int64
		expectedErr error
	}{
		{
			name:        "create task",
			task:        task,
			taskId:      0,
			expectedErr: nil,
		},
		{
			name:        "error if empty task",
			expectedErr: errors.New("empty name"),
			taskId:      -1,
			task:        empty_task,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        INSERT INTO
          tasks (name, status, projectId)
        VALUES ($1, $2, $3);`

			mock.
				ExpectExec(query).
				WithArgs(
					test.task.Name,
					test.task.Status,
					test.task.ProjectId,
				).
				WillReturnResult(sqlmock.NewResult(test.taskId, 1))

			repo := repository.NewTasksRepository(db)
			id, createErr := repo.CreateTask(test.task)

			assert.Equal(t, test.expectedErr, createErr)
			assert.Equal(t, test.taskId, id)
		})
	}
}
