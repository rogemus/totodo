package pkg_test

import (
	"testing"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateList(t *testing.T) {
	list_original := model.List{
		Id:   1,
		Name: "test list",
	}
	list_updated := model.List{
		Id:   1,
		Name: "udpated",
	}
	testCases := []struct {
		name           string
		originalList   model.List
		updatedList    model.List
		expectedErr    error
		expectedSqlErr error
	}{
		{
			name:           "update task",
			originalList:   list_original,
			updatedList:    list_updated,
			expectedErr:    nil,
			expectedSqlErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        UPDATE 
          lists AS l
        SET
          l.name = $2
        WHERE
          l.id = $1;
      `
			mock.
				ExpectExec(query).
				WithArgs(test.updatedList.Name, test.originalList.Id).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewListRepository(db)
			updateErr := repo.UpdateList(test.updatedList)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, test.expectedErr, updateErr)
			assert.Equal(t, test.expectedSqlErr, sqlErr)
		})
	}
}
