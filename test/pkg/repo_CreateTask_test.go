package pkg_test

import (
	"errors"
	"testing"
	"time"
	"totodo/pkg"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateTask(t *testing.T) {
	var empty_task pkg.Task

	createdDate, _ := time.Parse("2006-01-02 15:04:05", "2024-09-08 19:15:17")
	task := pkg.Task{
		Id:          1,
		Description: "test task",
		Created:     createdDate,
	}

	testCases := []struct {
		name        string
		task        pkg.Task
		expectedErr error
	}{
		{
			name:        "create task",
			task:        task,
			expectedErr: nil,
		},
		{
			name:        "error if empty task",
			expectedErr: errors.New("empty description"),
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
				WillReturnResult(sqlmock.NewResult(int64(test.task.Id), 1))

			repo := pkg.NewTasksRepository(db)
			createErr := repo.CreateTask(test.task)

			assert.Equal(t, test.expectedErr, createErr)
		})
	}
}
