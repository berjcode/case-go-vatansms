package main

import (
	"berjcode/dependency/database"
	"berjcode/dependency/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		panic("Failed to connect to database")
	}
	defer db.Close()

	if err := database.Migrate(db); err != nil {
		panic("Failed to migrate database schema")
	}
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	staticHandler := handlers.NewStaticHandler()
	e.Static("/static", "static")

	e.GET("/", staticHandler.IndexHTML)
	e.GET("/register", staticHandler.RegisterHTML)
	e.GET("/plan", staticHandler.PlanPageHTML)
	e.POST("/users", handlers.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}
