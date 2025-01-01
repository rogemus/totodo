package pkg_test

import (
	"errors"
	"testing"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateList(t *testing.T) {
	var empty_list model.List

	list := model.List{
		Name: "test list",
	}

	testCases := []struct {
		name        string
		list        model.List
		listId      int64
		expectedErr error
	}{
		{
			name:        "create list",
			list:        list,
			listId:      0,
			expectedErr: nil,
		},
		{
			name:        "error if empty name",
			expectedErr: errors.New("empty name"),
			list:        empty_list,
			listId:      -1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        INSERT INTO
          lists (name)
        VALUES ($1);`

			mock.
				ExpectExec(query).
				WithArgs(
					test.list.Name,
				).
				WillReturnResult(sqlmock.NewResult(test.listId, 1))

			repo := repository.NewListRepository(db)
			id, createErr := repo.CreateList(test.list)

			assert.Equal(t, test.expectedErr, createErr)
			assert.Equal(t, test.listId, id)
		})
	}
}
