package helper

import "github.com/gofiber/fiber/v2"

type JSON struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	json := JSON{
		Message: message,
		Data:    data,
	}
	return c.Status(statusCode).JSON(json)
}
