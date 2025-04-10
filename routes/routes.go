package routes

import (
	"todo_api/config"
	"todo_api/controller"
	"todo_api/middleware"
	"todo_api/repositories"
	"todo_api/services"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	db := config.NewDb()

	userRepo := repositories.NewUserRepositoryImpl()
	userService := services.NewUserServiceImpl(userRepo, db)
	userController := controller.NewUserController(userService)

	e.POST("/register", userController.Register)
	e.POST("/login", userController.Login)

	listRepo := repositories.NewListRepositoryImpl()
	listService := services.NewListServiceImpl(listRepo, db)
	listController := controller.NewListControllerImpl(listService)

	e.POST("/list", listController.Create, middleware.JWTMiddleware)
	e.PUT("/list/:list_id", listController.Update, middleware.JWTMiddleware)
	e.DELETE("/list/:list_id", listController.Delete, middleware.JWTMiddleware)
	e.GET("/list", listController.FindByUserId, middleware.JWTMiddleware)
}
