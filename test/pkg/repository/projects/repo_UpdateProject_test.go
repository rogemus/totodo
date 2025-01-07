package pkg_test

import (
	"testing"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_UpdateProject(t *testing.T) {
	project_original := model.Project{
		Id:   1,
		Name: "test list",
	}
	project_updated := model.Project{
		Id:   1,
		Name: "udpated",
	}
	testCases := []struct {
		name            string
		originalProject model.Project
		updatedProject  model.Project
		expectedErr     error
		expectedSqlErr  error
	}{
		{
			name:            "update project",
			originalProject: project_original,
			updatedProject:  project_updated,
			expectedErr:     nil,
			expectedSqlErr:  nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			query := `
        UPDATE 
          projects AS p
        SET
          p.name = $2
        WHERE
          p.id = $1;
      `
			mock.
				ExpectExec(query).
				WithArgs(test.updatedProject.Name, test.originalProject.Id).
				WillReturnResult(sqlmock.NewResult(1, 1)).
				WillReturnError(test.expectedSqlErr)

			repo := repository.NewProjectsRepository(db)
			updateErr := repo.UpdateProject(test.updatedProject)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, test.expectedErr, updateErr)
			assert.Equal(t, test.expectedSqlErr, sqlErr)
		})
	}
}
