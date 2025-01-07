package pkg_test

import (
	"testing"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_DeleteProject(t *testing.T) {
	testCases := []struct {
		name           string
		projectId      int
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "delete project",
			projectId:      1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
		{
			name:           "no error if no project to delete",
			projectId:      -1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        DELETE FROM
          projects AS p
        WHERE
          p.id = $1;
      `
			mock.
				ExpectExec(query).
				WithArgs(test.projectId).
				WillReturnResult(sqlmock.NewResult(int64(test.projectId), 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewProjectsRepository(db)
			deleteErr := repo.DeleteProject(test.projectId)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, test.expectedErr, deleteErr)
			assert.Equal(t, test.expectedSqlErr, sqlErr)
		})
	}
}
