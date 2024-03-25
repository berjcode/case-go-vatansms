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
	// Login
	e.POST("/v1/auth", handlers.Login)

	// User
	e.GET("/v1/users/:id", handlers.GetUserData, mymiddleware.AuthMiddleware)
	e.PUT("/v1/users", handlers.UpdateUser, mymiddleware.AuthMiddleware)
	e.POST("/v1/users", handlers.CreateUser)

	// Lesson
	e.PUT("/v1/lessons", staticHandler.LessonDetailPageHtml, mymiddleware.AuthMiddleware)
	e.PUT("/v1/lessons", handlers.UpdateLesson, mymiddleware.AuthMiddleware)
	e.POST("/v1/lessons", handlers.CreateUserLesson, mymiddleware.AuthMiddleware)
	e.GET("/v1/lessons/:id", handlers.GetLessonById, mymiddleware.AuthMiddleware)
	e.GET("/v1/lessons/user/:userid", handlers.GetAllLessonsByUser, mymiddleware.AuthMiddleware)

	//Plan Status
	e.POST("/v1/planstatuses", handlers.CreatePlanStatus, mymiddleware.AuthMiddleware)
	e.PUT("/v1/planstatuses", handlers.UpdatePlanStatus, mymiddleware.AuthMiddleware)
	e.GET("/v1/planstatuses", handlers.GetAllPlanStatus, mymiddleware.AuthMiddleware)
	e.GET("/v1/planstatuses/:id", handlers.GetPlanStatusById, mymiddleware.AuthMiddleware)

	//Plan
	e.POST("/v1/plans", handlers.CreatePlan, mymiddleware.AuthMiddleware)
	e.PUT("/v1/plans", handlers.UpdatePlan, mymiddleware.AuthMiddleware)
	e.GET("/v1/plans/:id", handlers.GetPlanById, mymiddleware.AuthMiddleware)
	e.GET("/v1/plans/:userid", handlers.GetPlanDetails, mymiddleware.AuthMiddleware)

	// e.GET("/", staticHandler.IndexHTML)
	// e.GET("/login", staticHandler.LoginPageHtml)
	// e.GET("/plan", staticHandler.PlanPageHTML, mymiddleware.AuthMiddleware)
	// e.GET("/register", staticHandler.RegisterHTML)
	// e.GET("/lesson", staticHandler.LessonPageHtml, mymiddleware.AuthMiddleware)
	// e.GET("/userdetail", staticHandler.UserDetailPage, mymiddleware.AuthMiddleware)
	e.Logger.Fatal(e.Start(":8080"))
}
