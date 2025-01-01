package pkg_test

import (
	"testing"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_DeleteTask(t *testing.T) {
	testCases := []struct {
		name           string
		taskId         int
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "delete task",
			taskId:         1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
		{
			name:           "no error if no task to delete",
			taskId:         -1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        DELETE FROM
          tasks AS t
        WHERE
          t.id = $1;
      `
			mock.
				ExpectExec(query).
				WithArgs(test.taskId).
				WillReturnResult(sqlmock.NewResult(int64(test.taskId), 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewTasksRepository(db)
			deleteErr := repo.DeleteTask(test.taskId)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, test.expectedErr, deleteErr)
			assert.Equal(t, test.expectedSqlErr, sqlErr)
		})
	}
}
