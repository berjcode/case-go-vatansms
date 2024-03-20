package main

import (
	"berjcode/dependency/database"
	"berjcode/dependency/handlers"
	"berjcode/dependency/mymiddleware"

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
	database.Migrate(db)
	e.Static("/static", "static")
	e.GET("/", staticHandler.IndexHTML)
	e.GET("/register", staticHandler.RegisterHTML)
	e.GET("/plan", staticHandler.PlanPageHTML, mymiddleware.AuthMiddleware)
	e.GET("/userdetail", staticHandler.UserDetailPage, mymiddleware.AuthMiddleware)
	e.GET("/getuserdetail", staticHandler.UserDetailData)
	e.POST("/users", handlers.CreateUser)
	e.POST("/login", handlers.Login)
	e.GET("/login", staticHandler.LoginPageHtml)

	e.Logger.Fatal(e.Start(":8080"))
}
