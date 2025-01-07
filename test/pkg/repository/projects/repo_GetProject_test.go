package pkg_test

import (
	"database/sql"
	"testing"
	"time"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

func Test_GetProject(t *testing.T) {
	var empty_list model.Project
	createdDate, _ := time.Parse("2006-01-02 15:04:05", "2024-09-08 19:15:17")
	project := model.Project{
		Id:      1,
		Name:    "test project",
		Created: createdDate,
	}

	testCases := []struct {
		name        string
		expected    model.Project
		projectId   int
		expectedErr error
	}{
		{
			name:        "returns projects",
			expected:    project,
			expectedErr: nil,
		},
		{
			name:        "returns empty projects",
			expected:    empty_list,
			expectedErr: sql.ErrNoRows,
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

			if !cmp.Equal(test.expected, empty_list) {
				expectedRows.
					AddRow(
						test.expected.Id,
						test.expected.Name,
						test.expected.Created,
					)
			}

			query := `
        SELECT
          p.id,
          p.name,
          p.created
        FROM
          projects AS p
        WHERE
          p.id = $1;`

			mock.
				ExpectQuery(query).
				WithArgs(test.expected.Id).
				WillReturnRows(expectedRows)

			repo := repository.NewProjectsRepository(db)
			result, getErr := repo.GetProject(test.expected.Id)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, sqlErr, nil)
			assert.Equal(t, test.expectedErr, getErr)
			assert.Equal(t, test.expected, result)
		})
	}
}
