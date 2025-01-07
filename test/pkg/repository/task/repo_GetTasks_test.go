package pkg_test

import (
	"testing"
	"time"
	"totodo/pkg/model"
	"totodo/pkg/repository"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func Test_GetTasks(t *testing.T) {
	empty_tasks := make([]model.Task, 0)
	createdDate, _ := time.Parse("2006-01-02 15:04:05", "2024-09-08 19:15:17")
	task_1 := model.Task{
		Id:          1,
		Name:        "test task",
		Created:     createdDate,
		Status:      "todo",
		ProjectId:   1,
		ProjectName: "testing",
	}
	task_2 := model.Task{
		Id:          2,
		Name:        "test task 2",
		Created:     createdDate,
		Status:      "todo",
		ProjectId:   1,
		ProjectName: "testing",
	}
	tasks := append(empty_tasks, task_1)
	tasks = append(tasks, task_2)

	testCases := []struct {
		name        string
		expected    []model.Task
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
				"name",
				"created",
				"status",
				"projectId",
				"projectName",
			}
			expectedRows := sqlmock.NewRows(columns)

			for _, task := range test.expected {
				expectedRows.AddRow(
					task.Id,
					task.Name,
					task.Created,
					task.Status,
					task.ProjectId,
					task.ProjectName,
				)
			}

			query := `
        SELECT
          t.id,
          t.name,
          t.created,
          t.status,
          t.projectId,
          p.name AS projectName
        FROM
          tasks AS t LEFT OUTER JOIN projects as p
        ON
          t.projectId = p.id
        WHERE
          p.id = $1
        ORDER BY
          t.created
        DESC;`

			mock.
				ExpectQuery(query).
				WithArgs(0).
				WillReturnRows(expectedRows)

			defer db.Close()

			repo := repository.NewTasksRepository(db)
			result, getErr := repo.GetTasks(0)
			sqlErr := mock.ExpectationsWereMet()

			assert.Equal(t, sqlErr, nil)
			assert.Equal(t, test.expectedErr, getErr)
			assert.Equal(t, len(test.expected), len(result))
			assert.Equal(t, test.expected, result)
		})
	}
}
