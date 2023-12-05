package handler

import (
	"cc-course-service/db"
	"cc-course-service/helper"
	"cc-course-service/model"
	"cc-course-service/repo"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	//unofficial openai go client
)

type CourseHandler struct {
	Repo repo.CourseRepository
}

func NewCourseHandler(repo repo.CourseRepository) *CourseHandler {
	return &CourseHandler{
		Repo: repo,
	}
}

func (handler *CourseHandler) CreateCourseHandler(c *fiber.Ctx) error {
	tx := db.DB.Begin()
	var raw model.CourseRaw
	if err := c.BodyParser(&raw); err != nil {
		return helper.Response(c, 400, "Error Parsing the Body", nil)
	}

	token := c.Get("Authorization")
	token = token[len("Bearer "):]
	claims, err := helper.VerifyToken(token)
	if err != nil {
		return helper.Response(c, 401, "Unauthorized", nil)
	}

	var responseMap map[string]interface{}
	decoder := json.NewDecoder(strings.NewReader(raw.Course))
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&responseMap); err != nil {
		tx.Rollback()
		if terr, ok := err.(*json.UnmarshalTypeError); ok {
			errorMessage := fmt.Sprintf("Error: Field '%s' memiliki tipe data yang salah. Harap periksa kembali inputan Anda", terr.Field)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": errorMessage,
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	title := responseMap["title"].(string)
	desc := responseMap["desc"].(string)
	duration := responseMap["duration"].(string)
	theme := responseMap["theme_activity"].(string)
	typeA := responseMap["type_activity"].(string)

	courseID := uuid.New()
	course := model.Course{
		ID:        courseID,
		UserID:    claims.Id,
		Title:     title,
		Desc:      desc,
		Duration:  duration,
		Theme:     theme,
		Type:      typeA,
		IsDone:    false,
		CreatedAt: helper.GetCurrentTime(),
		Prompt:    raw.Prompt,
		Subtopic:  []model.Subtopic{},
	}
	if err := db.DB.Create(&course).Error; err != nil {
		tx.Rollback()
		return err
	}

	subtopics := responseMap["subtitles"].([]interface{})
	for _, subtopic := range subtopics {
		subtitle := subtopic.(map[string]interface{})
		topic := subtitle["topic"].(string)
		shortDesc := subtitle["shortdesc"].(string)

		subtopicID := uuid.New()

		subtopicInstance := model.Subtopic{
			ID:       subtopicID,
			Desc:     shortDesc,
			Topic:    topic,
			CourseID: courseID,
			IsDone:   false,
			Content:  model.Content{},
		}
		if err := db.DB.Create(&subtopicInstance).Error; err != nil {
			tx.Rollback()
			return err
		}

		content := subtitle["content"].(map[string]interface{})
		opening := content["opening"].(string)
		closing := content["closing"].(string)

		contentID := uuid.New()

		contentInstance := model.Content{
			ID:      contentID,
			Opening: opening,
			Closing: closing,
			SubID:   subtopicID,
			Step:    []model.Step{},
		}
		if err := db.DB.Create(&contentInstance).Error; err != nil {
			tx.Rollback()
			return err
		}

		steps := content["step"].([]interface{})
		for _, step := range steps {
			stepText := step.(string)

			stepID := uuid.New()

			stepInstance := model.Step{
				ID:        stepID,
				Step:      stepText,
				ContentID: contentID,
			}
			if err := db.DB.Create(&stepInstance).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return helper.Response(c, 200, "Course Successfully Generated",
		fiber.Map{
			"course_id": courseID})
}

func (handler *CourseHandler) GetAllCourseHandler(c *fiber.Ctx) error {
	courses, err := handler.Repo.GetAll()

	if err != nil {
		return helper.Response(c, 400, "Course not found", nil)
	}

	return helper.Response(c, 200, "Course found", courses)
}

func (handler *CourseHandler) GetCourseByIdHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	course, err := handler.Repo.GetCourseByID(id)

	if err != nil {
		return helper.Response(c, 400, "Course not found", nil)
	}

	return helper.Response(c, 200, "Course found", course)
}

func (handler *CourseHandler) MarkAsDoneHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	err := handler.Repo.MarkSubtopicAsDone(id)

	if err != nil {
		return helper.Response(c, 400, "Subtopic not found", nil)
	}

	return helper.Response(c, 200, "Subtopic marked as done", nil)
}

func (handler *CourseHandler) TestFirebase(c *fiber.Ctx) error {
	return helper.Response(c, 200, "Success", nil)
}
