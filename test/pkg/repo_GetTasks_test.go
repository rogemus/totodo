package pkg_test

import (
	"testing"
	"time"
	"totodo/pkg"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetTasks(t *testing.T) {
	empty_tasks := make([]pkg.Task, 0)
	createdDate, _ := time.Parse("2006-01-02 15:04:05", "2024-09-08 19:15:17")
	task_1 := pkg.Task{
		Id:          1,
		Description: "test task",
		Created:     createdDate,
	}
	task_2 := pkg.Task{
		Id:          1,
		Description: "test task",
		Created:     createdDate,
	}
	tasks := append(empty_tasks, task_1)
	tasks = append(tasks, task_2)

	testCases := []struct {
		name        string
		expected    []pkg.Task
		expectedErr error
	}{
		{
			name:        "returns array of tasks",
			expected:    tasks,
			expectedErr: nil,
		},
		{
			name:        "returns empty array of tasks",
			expected:    empty_tasks,
			expectedErr: nil,
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

			for _, task := range test.expected {
				expectedRows.AddRow(
					task.Id,
					task.Description,
					task.Created,
				)
			}

			mock.
				ExpectQuery("SELECT id, description, created FROM tasks;").
				WithoutArgs().
				WillReturnRows(expectedRows)

			defer db.Close()

			repo := pkg.NewTasksRepository(db)
			result, getErr := repo.GetTasks()
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, sqlErr, nil)
			assert.Equal(t, test.expectedErr, getErr)
			assert.Equal(t, len(test.expected), len(result))
			assert.Equal(t, test.expected, result)
		})
	}
}
