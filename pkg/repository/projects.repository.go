package repository

import (
	"database/sql"
	_ "embed"
	"errors"
	"fmt"
	"totodo/pkg/model"
	"totodo/pkg/utils"
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
	UpdateProject(project model.Project) error
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
	utils.Log.Info(fmt.Sprintf("Getting Project: Project [%d]", id))

	var project model.Project

	row := r.db.QueryRow(getProjectQuery, id)
	err := row.Scan(
		&project.Id,
		&project.Name,
		&project.Created,
	)

	if err != nil {
		utils.Log.Error(err)
		return project, err
	}

	utils.Log.Success(fmt.Sprintf("Got Project: Project [%d] (%s)", project.Id, project.Name))
	return project, nil
}

func (r *projectsRepository) GetProjects() ([]model.Project, error) {
	utils.Log.Info("Getting All Projects")
	rows, err := r.db.Query(getProjectsQuery)

	if err != nil {
		utils.Log.Error(err)
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
			utils.Log.Error(err)
			return nil, err
		}

		projects = append(projects, project)
	}

	if err := rows.Err(); err != nil {
		utils.Log.Error(err)
		return nil, err
	}

	utils.Log.Success(fmt.Sprintf("Got Projects: Projects Count [%d]", len(projects)))
	return projects, nil
}

func (r *projectsRepository) UpdateProject(project model.Project) error {
	utils.Log.Info(fmt.Sprintf("Updating Project: Project [%d]", project.Id))

	_, err := r.db.Exec(updateProjectQuery, project.Name, project.Id)

	if err != nil {
		utils.Log.Error(err)
		return err
	}

	utils.Log.Success(fmt.Sprintf("Updated Project: Project [%d] (%s)", project.Id, project.Name))
	return nil
}

func (r *projectsRepository) DeleteProject(id int) error {
	utils.Log.Info(fmt.Sprintf("Deleting Project: Project [%d]", id))

	_, err := r.db.Exec(deleteProjectQuery, id)

	if err != nil {
		utils.Log.Error(err)
		return err
	}

	utils.Log.Success(fmt.Sprintf("Deleted Project: Project [%d]", id))
	return nil
}

func (r *projectsRepository) CreateProject(project model.Project) (int64, error) {
	utils.Log.Info(fmt.Sprintf("Creating Project: Project (%s)", project.Name))

	if project.Name == "" {
		return -1, errors.New("empty name")
	}

	result, err := r.db.Exec(createProjectQuery, project.Name)

	if err != nil {
		utils.Log.Error(err)
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		utils.Log.Error(err)
		return -1, err
	}

	utils.Log.Success(fmt.Sprintf("Project Created: Project [%d] (%s)", id, project.Name))
	return id, nil
}
