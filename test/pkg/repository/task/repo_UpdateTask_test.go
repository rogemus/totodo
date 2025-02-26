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
		Id:   1,
		Name: "test task",
	}
	task_updated := model.Task{
		Id:   1,
		Name: "udpated",
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

			query := `
        UPDATE 
          tasks AS t
        SET
          t.name = $2
        WHERE
          t.id = $1;
      `
			mock.
				ExpectExec(query).
				WithArgs(test.updatedTask.Name, test.originalTask.Id).
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
