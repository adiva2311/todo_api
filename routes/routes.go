package routes

import (
	"database/sql"
	"todo_api/controller"
	"todo_api/repositories"
	"todo_api/services"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo)  {
	var db *sql.DB
	userRepo := repositories.NewUserRepositoryImpl()
	userService := services.NewUserServiceImpl(userRepo, db)
	userController := controller.NewUserController(userService)

	e.POST("/register", userController.Register)
}