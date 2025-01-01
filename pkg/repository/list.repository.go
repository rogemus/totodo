package repository

import (
	"database/sql"
	_ "embed"
	"errors"
	"totodo/pkg/model"
)

var (
	//go:embed queries/lists/getList.sql
	getListQuery string
	//go:embed queries/lists/getLists.sql
	getListsQuery string
	//go:embed queries/lists/updateList.sql
	updateListQuery string
	//go:embed queries/lists/deleteList.sql
	deleteListQuery string
	//go:embed queries/lists/createList.sql
	createListQuery string
)

type ListRepository interface {
	GetList(id int) (model.List, error)
	GetLists() ([]model.List, error)
	UpdateList(list model.List) error
	DeleteList(id int) error
	CreateList(list model.List) (int64, error)
}

type listRepository struct {
	db *sql.DB
}

func NewListRepository(db *sql.DB) ListRepository {
	return &listRepository{db}
}

func (r *listRepository) GetList(id int) (model.List, error) {
	var list model.List

	row := r.db.QueryRow(getListQuery, id)
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

func (r *listRepository) GetLists() ([]model.List, error) {
	rows, err := r.db.Query(getListsQuery)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	lists := make([]model.List, 0)

	for rows.Next() {
		var list model.List

		err := rows.Scan(
			&list.Id,
			&list.Name,
			&list.Created,
		)

		if err != nil {
			return nil, err
		}

		lists = append(lists, list)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lists, nil
}

func (r *listRepository) UpdateList(list model.List) error {
	_, err := r.db.Exec(updateListQuery, list.Name, list.Id)

	if err != nil {
		return err
	}

	return nil
}

func (r *listRepository) DeleteList(id int) error {
	_, err := r.db.Exec(deleteListQuery, id)

	if err != nil {
		return err
	}

	return nil
}

func (r *listRepository) CreateList(list model.List) (int64, error) {
	if list.Name == "" {
		return -1, errors.New("empty name")
	}

	result, err := r.db.Exec(createListQuery, list.Name)

	if err != nil {
		return -1, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return -1, err
	}

	return id, nil
}
