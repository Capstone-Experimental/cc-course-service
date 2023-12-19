package handler

import (
	"cc-course-service/helper"
	"cc-course-service/repo"

	"github.com/gofiber/fiber/v2"
)

type DashboardHandler struct {
	Repo repo.DashboardRepository
}

func NewDashboardHandler(repo repo.DashboardRepository) *DashboardHandler {
	return &DashboardHandler{
		Repo: repo,
	}
}

func (handler *DashboardHandler) GetDashboard(c *fiber.Ctx) error {
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

	dashboard, err := handler.Repo.GetDashboard(userID)
	if err != nil {
		return helper.Response(c, 400, "Dashboard not found", nil)
	}

	return helper.Response(c, 200, "Dashboard found", dashboard)
}
