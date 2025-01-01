package pkg_test

import (
	"testing"
	"time"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetLists(t *testing.T) {
	empty_list := make([]model.List, 0)
	createdDate, _ := time.Parse("2006-01-02 15:04:05", "2024-09-08 19:15:17")
	list_1 := model.List{
		Id:      1,
		Name:    "test list",
		Created: createdDate,
	}
	list_2 := model.List{
		Id:      2,
		Name:    "test list 2",
		Created: createdDate,
	}
	lists := append(empty_list, list_1)
	lists = append(lists, list_2)

	testCases := []struct {
		name        string
		expected    []model.List
		expectedErr error
	}{
		{
			name:        "returns array of lists",
			expected:    lists,
			expectedErr: nil,
		},
		{
			name:        "returns empty array of lists",
			expected:    empty_list,
			expectedErr: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			columns := []string{
				"id",
				"name",
				"created",
			}
			expectedRows := sqlmock.NewRows(columns)

			for _, task := range test.expected {
				expectedRows.AddRow(
					task.Id,
					task.Name,
					task.Created,
				)
			}

			query := `
        SELECT
          l.id,
          l.name,
          l.created
        FROM
          lists AS l
        ORDER BY
          l.created
        DESC;`

			mock.
				ExpectQuery(query).
				WithoutArgs().
				WillReturnRows(expectedRows)

			defer db.Close()

			repo := repository.NewListRepository(db)
			result, getErr := repo.GetLists()
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, sqlErr, nil)
			assert.Equal(t, test.expectedErr, getErr)
			assert.Equal(t, len(test.expected), len(result))
			assert.Equal(t, test.expected, result)
		})
	}
}
