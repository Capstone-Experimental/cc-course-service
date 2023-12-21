package route

import (
	"cc-course-service/handler"
	"cc-course-service/repo"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, courseRepo repo.CourseRepository, feedbackRepo repo.FeedbackRepository, dashboardRepo repo.DashboardRepository) {
	courseHandler := handler.NewCourseHandler(courseRepo)

	courseRoutes := app.Group("/api/v1")

	courseRoutes.Get("/", courseHandler.GetAllCourseHandler)            // GET /course
	courseRoutes.Get("/my-course", courseHandler.MyCourseHandler)       // GET /course/my-course
	courseRoutes.Get("/course/:id", courseHandler.GetCourseByIdHandler) // GET /course/:id
	courseRoutes.Post("/create", courseHandler.CreateCourseHandler)     // POST /course
	courseRoutes.Put("/subtitle/:id", courseHandler.MarkAsDoneHandler)  // PUT /course/subtitle/:id

	feedbackHandler := handler.NewFeedbackHandler(feedbackRepo)

	feedbackRoutes := app.Group("/api/v1/feedback")

	feedbackRoutes.Post("/:course_id", feedbackHandler.CreateFeedbackHandler) // POST /feedback/:course_id

	dashboardHandler := handler.NewDashboardHandler(dashboardRepo)

	dashboardRoutes := app.Group("/api/v1/dashboard") // GET /dashboard

	dashboardRoutes.Get("/", dashboardHandler.GetDashboard)
}
