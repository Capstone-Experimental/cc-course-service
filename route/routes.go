package route

import (
	"cc-course-service/handler"
	"cc-course-service/repo"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, courseRepo repo.CourseRepository, feedbackRepo repo.FeedbackRepository, dashboardRepo repo.DashboardRepository) {
	courseHandler := handler.NewCourseHandler(courseRepo)

	courseRoutes := app.Group("/api/v1")

	courseRoutes.Get("/", courseHandler.GetAllCourseHandler)
	courseRoutes.Get("/my-course", courseHandler.MyCourseHandler)
	courseRoutes.Get("/course/:id", courseHandler.GetCourseByIdHandler)
	courseRoutes.Post("/create", courseHandler.CreateCourseHandler)
	courseRoutes.Put("/subtitle/:id", courseHandler.MarkAsDoneHandler)

	feedbackHandler := handler.NewFeedbackHandler(feedbackRepo)

	feedbackRoutes := app.Group("/api/v1/feedback")

	feedbackRoutes.Post("/:course_id", feedbackHandler.CreateFeedbackHandler)

	dashboardHandler := handler.NewDashboardHandler(dashboardRepo)

	dashboardRoutes := app.Group("/api/v1/dashboard")

	dashboardRoutes.Get("/", dashboardHandler.GetDashboard)
}
