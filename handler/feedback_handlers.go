package handler

import (
	"cc-course-service/db"
	"cc-course-service/helper"
	"cc-course-service/model"
	"cc-course-service/repo"
	"strconv"

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

// CreateFeedbackHandler handles POST /feedback/:course_id
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

	var course model.Course
	tx.Where("id = ?", courseId).First(&course)

	feedback := model.Feedback{
		ID:       uuid.New(),
		UserId:   userID,
		CourseId: courseId,
		Prompt:   course.Prompt,
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

// GetFeedbacksHandler handles GET /feedback
func (handler *FeedbackHandler) GetFeedbacks(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 0
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil || pageSize < 1 {
		pageSize = 0
	}

	feedbacks, err := handler.Repo.GetAllFeedbacks(page, pageSize)
	if err != nil {
		return helper.Response(c, 400, "Feedbacks not found", nil)
	}

	if page > 0 && pageSize > 0 {
		paginated, err := helper.FeedbackPaginate(c, page, pageSize, feedbacks)
		if err != nil {
			return helper.Response(c, 400, "Feedbacks not found", nil)
		}

		return helper.Response(c, 200, "Course found", paginated)
	}

	return helper.Response(c, 200, "Feedback found", feedbacks)
}
