package repository

import (
	"database/sql"
	_ "embed"
	"errors"
	"totodo/pkg/model"
)

var (
	//go:embed queries/projects/getProject.sql
	getProjectQuery string
	//go:embed queries/projects/getProjects.sql
	getProjectsQuery string
	//go:embed queries/projects/updateProject.sql
	updateProjectQuery string
	//go:embed queries/projects/deleteProject.sql
	deleteProjectQuery string
	//go:embed queries/projects/createProject.sql
	createProjectQuery string
)

type ProjectsRepository interface {
	GetProject(id int) (model.Project, error)
	GetProjects() ([]model.Project, error)
	UpdateProject(list model.Project) error
	DeleteProject(id int) error
	CreateProject(project model.Project) (int64, error)
}

type projectsRepository struct {
	db *sql.DB
}

func NewProjectsRepository(db *sql.DB) ProjectsRepository {
	return &projectsRepository{db}
}

func (r *projectsRepository) GetProject(id int) (model.Project, error) {
	var list model.Project

	row := r.db.QueryRow(getProjectQuery, id)
	err := row.Scan(
		&list.Id,
		&list.Name,
		&list.Created,
	)

	if err != nil {
		return list, err
	}

	return list, nil
}

func (r *projectsRepository) GetProjects() ([]model.Project, error) {
	rows, err := r.db.Query(getProjectsQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	projects := make([]model.Project, 0)

	for rows.Next() {
		var project model.Project

		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Created,
		)

		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectsRepository) UpdateProject(list model.Project) error {
	_, err := r.db.Exec(updateProjectQuery, list.Name, list.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *projectsRepository) DeleteProject(id int) error {
	_, err := r.db.Exec(deleteProjectQuery, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *projectsRepository) CreateProject(list model.Project) (int64, error) {
	if list.Name == "" {
		return -1, errors.New("empty name")
	}

	result, err := r.db.Exec(createProjectQuery, list.Name)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}
