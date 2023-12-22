package main

import (
	"example/web-service-gin/db"
	"example/web-service-gin/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	db.Init() // Initialize the database

	e := echo.New()

	// Define routes
	e.GET("/", handlers.Hello)
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.POST("/user", handlers.CreateUser)

	e.Logger.Fatal(e.Start(":1323"))
}
