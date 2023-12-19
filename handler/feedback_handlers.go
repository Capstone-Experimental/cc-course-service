package handler

import (
	"cc-course-service/db"
	"cc-course-service/helper"
	"cc-course-service/model"
	"cc-course-service/repo"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FeedbackHandler struct {
	Repo repo.FeedbackRepository
}

func NewFeedbackHandler(repo repo.FeedbackRepository) *FeedbackHandler {
	return &FeedbackHandler{
		Repo: repo,
	}
}

func (handler *FeedbackHandler) CreateFeedbackHandler(c *fiber.Ctx) error {
	tx := db.DB.Begin()

	courseId := c.Params("course_id")

	// JWT auth
	// token := c.Get("Authorization")
	// token = token[len("Bearer "):]
	// claims, err := helper.VerifyToken(token)
	// if err != nil {
	// 	return helper.Response(c, 401, "Unauthorized", nil)
	// }
	// var userID = claims.Id

	// Firebase auth
	claims, ok := c.Locals("claims").(map[string]interface{})
	if !ok {
		return helper.Response(c, 401, "Unauthorized", nil)
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return helper.Response(c, 500, "Internal Server Error", nil)
	}

	var raw model.FeedbackRaw
	if err := c.BodyParser(&raw); err != nil {
		return helper.Response(c, 400, "Error Parsing the Body", nil)
	}

	if raw.Rating > 5 || raw.Rating < 1 {
		return helper.Response(c, 400, "Rating must be between 1 and 5", nil)
	}

	feedback := model.Feedback{
		ID:       uuid.New(),
		UserId:   userID,
		CourseId: courseId,
		Feedback: raw.Feedback,
		Rating:   raw.Rating,
	}

	if err := db.DB.Create(&feedback).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !ok {
		return helper.Response(c, 500, "Internal Server Error", nil)
	}

	return helper.Response(c, 201, "Success", feedback)
}
