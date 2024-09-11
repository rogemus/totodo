package pkg_test

import (
	"errors"
	"testing"
	"totodo/pkg"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateTask(t *testing.T) {
	var empty_task pkg.Task

	task := pkg.Task{
		Description: "test task",
	}

	testCases := []struct {
		name        string
		task        pkg.Task
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
			expectedErr: errors.New("empty description"),
			taskId:      -1,
			task:        empty_task,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			mock.
				ExpectExec("INSERT INTO tasks (description) VALUES ($1)").
				WithArgs(
					test.task.Description,
				).
				WillReturnResult(sqlmock.NewResult(test.taskId, 1))

			repo := pkg.NewTasksRepository(db)
			id, createErr := repo.CreateTask(test.task)

			assert.Equal(t, test.expectedErr, createErr)
			assert.Equal(t, test.taskId, id)
		})
	}
}
