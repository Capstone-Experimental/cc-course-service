package helper

import (
	"cc-course-service/model"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// CoursePaginate is a helper function for pagination
func CoursePaginate(c *fiber.Ctx, page, pageSize int, items []model.Course) (map[string]interface{}, error) {

	var nextURL, prevURL string

	if page > 1 {
		prevURL = fmt.Sprintf("%s?page=%d&page_size=%d", c.OriginalURL(), page-1, pageSize)
	}

	if len(items) == pageSize {
		nextURL = fmt.Sprintf("%s?page=%d&page_size=%d", c.OriginalURL(), page+1, pageSize)
	}

	response := map[string]interface{}{
		"count":    len(items),
		"next":     nextURL,
		"previous": prevURL,
		"results":  items}

	return response, nil
}

// FeedbackPaginate is a helper function for pagination
func FeedbackPaginate(c *fiber.Ctx, page, pageSize int, items []model.Feedback) (map[string]interface{}, error) {

	var nextURL, prevURL string

	if page > 1 {
		prevURL = fmt.Sprintf("%s?page=%d&page_size=%d", c.OriginalURL(), page-1, pageSize)
	}

	if len(items) == pageSize {
		nextURL = fmt.Sprintf("%s?page=%d&page_size=%d", c.OriginalURL(), page+1, pageSize)
	}

	response := map[string]interface{}{
		"count":    len(items),
		"next":     nextURL,
		"previous": prevURL,
		"results":  items}

	return response, nil
}
