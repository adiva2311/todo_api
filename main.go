package main

import (
	"todo_api/config"
	"todo_api/routes"

	//_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.NewDb()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start("127.0.0.1:8080"))
}