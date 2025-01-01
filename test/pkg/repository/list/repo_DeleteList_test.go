package pkg_test

import (
	"testing"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_DeleteList(t *testing.T) {
	testCases := []struct {
		name           string
		listId         int
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "delete list",
			listId:         1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
		{
			name:           "no error if no list to delete",
			listId:         -1,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        DELETE FROM
          lists AS l
        WHERE
          l.id = $1;
      `
			mock.
				ExpectExec(query).
				WithArgs(test.listId).
				WillReturnResult(sqlmock.NewResult(int64(test.listId), 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewListRepository(db)
			deleteErr := repo.DeleteList(test.listId)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, test.expectedErr, deleteErr)
			assert.Equal(t, test.expectedSqlErr, sqlErr)
		})
	}
}
