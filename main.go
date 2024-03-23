package main

import (
	"berjcode/dependency/database"
	"berjcode/dependency/handlers"
	"berjcode/dependency/mymiddleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := database.NewDB("dbconfig.json")
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
	e.POST("/login", handlers.Login)
	e.POST("/users", handlers.CreateUser)
	e.GET("/lessons", handlers.GetAllLessonsByUser)
	e.GET("/lesson/:id", handlers.GetLessonById, mymiddleware.AuthMiddleware)
	e.PUT("/updateusers", handlers.UpdateUser, mymiddleware.AuthMiddleware)
	e.POST("/lesson", handlers.CreateUserLesson, mymiddleware.AuthMiddleware)
	e.PUT("/updatelesson", handlers.UpdateLesson, mymiddleware.AuthMiddleware)
	e.PUT("/lesson", staticHandler.LessonDetailPageHtml, mymiddleware.AuthMiddleware)
	e.GET("/users/:id", handlers.GetUserData, mymiddleware.AuthMiddleware)

	e.GET("/", staticHandler.IndexHTML)
	e.GET("/login", staticHandler.LoginPageHtml)
	e.GET("/plan", staticHandler.PlanPageHTML, mymiddleware.AuthMiddleware)
	e.GET("/register", staticHandler.RegisterHTML)
	e.GET("/lesson", staticHandler.LessonPageHtml, mymiddleware.AuthMiddleware)
	e.GET("/userdetail", staticHandler.UserDetailPage, mymiddleware.AuthMiddleware)
	e.Logger.Fatal(e.Start(":8080"))
}
