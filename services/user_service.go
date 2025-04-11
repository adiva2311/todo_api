package services

import (
	"database/sql"
	"errors"
	"todo_api/helper"
	"todo_api/models"
	"todo_api/repositories"

	"github.com/labstack/echo/v4"
)

type UserService interface {
	Register(c echo.Context, request models.User) (helper.UserResponse, error)
	Login(c echo.Context, request helper.LoginRequest) (helper.LoginResponse, error)
}

type UserServiceImpl struct {
	UserRepo repositories.UserRepository
	Db       *sql.DB
}

// Register implements UserService.
func (service *UserServiceImpl) Register(c echo.Context, request models.User) (helper.UserResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	hashPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		panic(err)
	}

	user := &models.User{
		Name:     request.Name,
		Username: request.Username,
		Email:    request.Email,
		Password: hashPassword,
	}

	savedUser, err := service.UserRepo.Register(c, tx, *user)
	if err != nil {
		panic(err)
	}

	return helper.ToRegisterResponse(savedUser), nil
}

// Login implements UserService.
func (service *UserServiceImpl) Login(c echo.Context, request helper.LoginRequest) (helper.LoginResponse, error) {
	tx, err := service.Db.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	//Check Username
	user, err := service.UserRepo.FindByUsername(c, tx, request.Username)
	if err != nil {
		return helper.LoginResponse{}, errors.New("invalid username or password")
	}

	//Check Password
	if !helper.CheckPasswordHash(request.Password, user.Password) {
		return helper.LoginResponse{}, errors.New("invalid username or password")
	}

	token, err := helper.GenerateJWT(int(user.Id))
	if err != nil {
		return helper.LoginResponse{}, err
	}
	return helper.ToLoginResponse(user, token), nil
}

func NewUserServiceImpl(userRepo repositories.UserRepository, db *sql.DB) UserService {
	return &UserServiceImpl{
		UserRepo: userRepo,
		Db:       db,
	}
}
