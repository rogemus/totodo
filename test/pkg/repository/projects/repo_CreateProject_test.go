package pkg_test

import (
	"errors"
	"testing"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_CreateProject(t *testing.T) {
	var empty_list model.Project

	project := model.Project{
		Name: "test list",
	}

	testCases := []struct {
		name        string
		project     model.Project
		projectId   int64
		expectedErr error
	}{
		{
			name:        "create project",
			project:     project,
			projectId:   0,
			expectedErr: nil,
		},
		{
			name:        "error if empty name",
			expectedErr: errors.New("empty name"),
			project:     empty_list,
			projectId:   -1,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        INSERT INTO
          projects (name)
        VALUES ($1);`

			mock.
				ExpectExec(query).
				WithArgs(
					test.project.Name,
				).
				WillReturnResult(sqlmock.NewResult(test.projectId, 1))

			repo := repository.NewProjectsRepository(db)
			id, createErr := repo.CreateProject(test.project)

			assert.Equal(t, test.expectedErr, createErr)
			assert.Equal(t, test.projectId, id)
		})
	}
}
