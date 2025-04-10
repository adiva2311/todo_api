package controller

import (
	"net/http"
	"todo_api/models"
	"todo_api/services"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	Register(c echo.Context) error
}

type UserControllerImpl struct {
	UserService services.UserService
}

// Register implements UserController.
func (controller *UserControllerImpl) Register(c echo.Context) error {
	//user := models.User{}
	userPayload := new(models.User)

	err := c.Bind(userPayload)
	if err != nil{
		panic(err)
	}

	result, err := controller.UserService.Register(c, models.User{
		Name: userPayload.Name,
		Username: userPayload.Username,
		Email: userPayload.Email,
		Password: userPayload.Password,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"message": "Gagal register", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}
