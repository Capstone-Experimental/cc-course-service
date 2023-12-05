package middleware

import (
	"cc-course-service/helper"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		header := c.Get("Authorization")
		if header == "" {
			return helper.Response(c, 401, "Unauthorized", nil)
		}
		parts := strings.Split(header, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return helper.Response(c, 401, "Unauthorized", nil)
		}
		token := parts[1]

		claims, err := helper.VerifyToken(token)
		if err != nil {
			return helper.Response(c, 401, "Unauthorized", nil)
		}
		log.Println("[Username]" + claims.Username)
		log.Println("[Id]" + claims.Id)

		return c.Next()
	}
}
