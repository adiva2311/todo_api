package repositories

import (
	"database/sql"
	"todo_api/models"

	"github.com/labstack/echo/v4"
)

type ListRepository interface {
	Create(c echo.Context, tx *sql.Tx, list models.List) (models.List, error)
	Update(c echo.Context, tx *sql.Tx, list models.List) (models.List, error)
	Delete(c echo.Context, tx *sql.Tx, listId uint)
	FindByUserId(c echo.Context, tx *sql.Tx, userId int) ([]models.List, error)
	FindId(c echo.Context, tx *sql.Tx, userId int) models.List
}

type ListRepositoryImpl struct {
}

// Create implements ListRepository.
func (repo *ListRepositoryImpl) Create(c echo.Context, tx *sql.Tx, list models.List) (models.List, error) {
	query := "INSERT INTO list (title, information, completed, user_id) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(c.Request().Context(), query, list.Title, list.Information, list.Completed, list.UserId)
	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	list.Id = uint(id)
	return list, nil
}

// Update implements ListRepository.
func (repo *ListRepositoryImpl) Update(c echo.Context, tx *sql.Tx, list models.List) (models.List, error) {
	query := "UPDATE list SET title = ?, information = ?, completed = ? WHERE id = ?"
	result, err := tx.ExecContext(c.Request().Context(), query, list.Title, list.Information, list.Completed, list.Id)
	if err != nil {
		panic(err)
	}

	id, err := result.RowsAffected()
	if id < 1 {
		panic(err)
	}

	return list, nil
}

// Delete implements ListRepository.
func (repo *ListRepositoryImpl) Delete(c echo.Context, tx *sql.Tx, listId uint) {
	query := "DELETE FROM list WHERE id = ?"
	result, err := tx.ExecContext(c.Request().Context(), query, listId)
	if err != nil {
		panic(err)
	}

	id, err := result.RowsAffected()
	if id < 1 {
		panic(err)
	}
}

// FindByUserId implements ListRepository.
func (repo *ListRepositoryImpl) FindByUserId(c echo.Context, tx *sql.Tx, userId int) ([]models.List, error) {
	query := "SELECT id, title, information, completed, user_id FROM list WHERE user_id = ?"
	rows, err := tx.QueryContext(c.Request().Context(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lists []models.List
	for rows.Next() {
		var list models.List
		err := rows.Scan(&list.Id, &list.Title, &list.Information, &list.Completed, &list.UserId)
		if err != nil {
			panic(err)
		}
		lists = append(lists, list)
	}
	return lists, nil
}

// FindId implements ListRepository.
func (repo *ListRepositoryImpl) FindId(c echo.Context, tx *sql.Tx, userId int) models.List {
	query := "SELECT id, title, information, completed, user_id WHERE id = ?"
	rows := tx.QueryRowContext(c.Request().Context(), query, userId)
	
	list := models.List{}
	err := rows.Scan(&list.Id, &list.Title, &list.Information, &list.Completed, &list.UserId)
	if err != nil {
		return list
	}
	return list
}

func NewListControllerImpl() ListRepository {
	return &ListRepositoryImpl{}
}
