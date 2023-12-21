package main

import (
	"cc-course-service/db"
	"cc-course-service/middleware"
	"cc-course-service/repo"
	"cc-course-service/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.InitDatabase()

	courseRepo := repo.NewCourseRepository(db.DB)
	feedbackRepo := repo.NewFeedbackRepository(db.DB)
	dashboardRepo := repo.NewDashboardRepository(db.DB)

	// Auth middleware uncomment one of these, you might need to do the same in handler and/or repo
	// app.Use(middleware.JWTProtected())
	app.Use(middleware.FirebaseAuth())

	route.InitRoutes(app, *courseRepo, *feedbackRepo, *dashboardRepo)

	app.Listen(":8081")
}
