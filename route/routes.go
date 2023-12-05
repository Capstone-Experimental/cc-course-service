package route

import (
	"cc-course-service/handler"
	"cc-course-service/repo"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, courseRepo repo.CourseRepository, feedbackRepo repo.FeedbackRepository) {
	courseHandler := handler.NewCourseHandler(courseRepo)

	courseRoutes := app.Group("/api/v1/course")

	courseRoutes.Get("/", courseHandler.GetAllCourseHandler)
	courseRoutes.Get("/:id", courseHandler.GetCourseByIdHandler)
	courseRoutes.Post("/create", courseHandler.CreateCourseHandler)
	courseRoutes.Put("/:id", courseHandler.MarkAsDoneHandler)

	feedbackHandler := handler.NewFeedbackHandler(feedbackRepo)

	feedbackRoutes := app.Group("/api/v1/feedback")

	feedbackRoutes.Post("/:course_id", feedbackHandler.CreateFeedbackHandler)
}
