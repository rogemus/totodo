package pkg_test

import (
	"testing"
	"time"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetProjects(t *testing.T) {
	empty_list := make([]model.Project, 0)
	createdDate, _ := time.Parse("2006-01-02 15:04:05", "2024-09-08 19:15:17")
	project_1 := model.Project{
		Id:      1,
		Name:    "test project",
		Created: createdDate,
	}
	project_2 := model.Project{
		Id:      2,
		Name:    "test project 2",
		Created: createdDate,
	}
	projects := append(empty_list, project_1)
	projects = append(projects, project_2)

	testCases := []struct {
		name        string
		expected    []model.Project
		expectedErr error
	}{
		{
			name:        "returns array of projects",
			expected:    projects,
			expectedErr: nil,
		},
		{
			name:        "returns empty array of projects",
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
          p.id,
          p.name,
          p.created
        FROM
          projects AS p
        ORDER BY
          p.created
        DESC;`

			mock.
				ExpectQuery(query).
				WithoutArgs().
				WillReturnRows(expectedRows)

			defer db.Close()

			repo := repository.NewProjectsRepository(db)
			result, getErr := repo.GetProjects()
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, sqlErr, nil)
			assert.Equal(t, test.expectedErr, getErr)
			assert.Equal(t, len(test.expected), len(result))
			assert.Equal(t, test.expected, result)
		})
	}
}
