package main

import (
	"todo_api/config"
	"todo_api/controller"
	"todo_api/repositories"
	"todo_api/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	db := config.NewDb()

	userRepo := repositories.NewUserRepositoryImpl()
	userService := services.NewUserServiceImpl(userRepo, db)
	userController := controller.NewUserController(userService)

	e.POST("/register", userController.Register)

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}