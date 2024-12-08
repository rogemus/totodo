package pkg_test

import (
	"testing"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateTask(t *testing.T) {
	task_original := model.Task{
		Id:          1,
		Description: "test task",
	}
	task_updated := model.Task{
		Id:          1,
		Description: "udpated",
	}
	testCases := []struct {
		name           string
		originalTask   model.Task
		updatedTask    model.Task
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "update task",
			originalTask:   task_original,
			updatedTask:    task_updated,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			mock.
				ExpectExec("UPDATE tasks SET description=$2 WHERE id = $1;").
				WithArgs(test.originalTask.Id, test.updatedTask.Description).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewTasksRepository(db)
			updateErr := repo.UpdateTask(test.updatedTask)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, test.expectedErr, updateErr)
			assert.Equal(t, test.expectedSqlErr, sqlErr)
		})
	}
}
