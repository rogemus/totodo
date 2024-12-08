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

func Test_GetTask(t *testing.T) {
	var empty_task model.Task
	createdDate, _ := time.Parse("2006-01-02 15:04:05", "2024-09-08 19:15:17")
	task := model.Task{
		Id:          1,
		Description: "test task",
		Created:     createdDate,
	}

	testCases := []struct {
		name        string
		expected    model.Task
		taskId      int
		expectedErr error
	}{
		{
			name:        "returns task",
			expected:    task,
			expectedErr: nil,
		},
		{
			name:        "returns empty task",
			expected:    empty_task,
			expectedErr: sql.ErrNoRows,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

			columns := []string{
				"id",
				"description",
				"created",
			}
			expectedRows := sqlmock.NewRows(columns)

			if !cmp.Equal(test.expected, empty_task) {
				expectedRows.
					AddRow(
						test.expected.Id,
						test.expected.Description,
						test.expected.Created,
					)
			}

			mock.
				ExpectQuery("SELECT id, description, created FROM tasks WHERE id = $1;").
				WithArgs(test.expected.Id).
				WillReturnRows(expectedRows)

			repo := repository.NewTasksRepository(db)
			result, getErr := repo.GetTask(test.expected.Id)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, sqlErr, nil)
			assert.Equal(t, test.expectedErr, getErr)
			assert.Equal(t, test.expected, result)
		})
	}
}
