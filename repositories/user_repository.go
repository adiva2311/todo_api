package repositories

import (
	"database/sql"
	"todo_api/models"

	"github.com/labstack/echo/v4"
)

type UserRepositoryImpl struct {
	//db *sql.DB
}

// Create implements UserRepository.
func (repo *UserRepositoryImpl) Register(c echo.Context, tx *sql.Tx, user models.User) (models.User, error) {
	query := "INSERT INTO user (name, username, email, password) VALUES (?, ?, ?, ?)"
	result, err := tx.ExecContext(c.Request().Context(), query, user.Name, user.Username, user.Email, user.Password)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	user.Id = uint(id)
	return user, nil
}

// FindById implements UserRepository.
func (repo *UserRepositoryImpl) FindByUsername(c echo.Context, tx *sql.Tx, username string) (models.User, error) {
	query := "SELECT name, email, password FROM user WHERE username = ?"
	rows := tx.QueryRowContext(c.Request().Context(), query, username)

	user := models.User{}
	err := rows.Scan(&user.Name, &user.Username, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}


type UserRepository interface {
	Register(c echo.Context, tx *sql.Tx, user models.User) (models.User, error)
	FindByUsername(c echo.Context, tx *sql.Tx, username string) (models.User, error)
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}
