package services

import (
	"database/sql"
	"todo_api/helper"
	"todo_api/models"
	"todo_api/repositories"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	Register(c echo.Context, request models.User) (helper.UserResponse, error)
}

type UserServiceImpl struct {
	UserRepo repositories.UserRepository
	Db       *sql.DB
}

// Register implements UserService.
func (service *UserServiceImpl) Register(c echo.Context, request models.User) (helper.UserResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil{
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	hashPassword, err := helper.HashPassword(request.Password)
	if err != nil{
		panic(err)
	}

	user := &models.User{
		Name: request.Name,
		Username: request.Username,
		Email: request.Email,
		Password: hashPassword,
	}

	savedUser, err := service.UserRepo.Register(c, tx, *user)
	if err != nil{
		panic(err)
	}

	return helper.ToRegisterResponse(savedUser), nil
}

func NewUserServiceImpl(userRepo repositories.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		Db:       db,
	}
}
