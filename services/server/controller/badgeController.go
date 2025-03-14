package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	// "github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetBadges(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetBadges endpoint called")
	return response.Success(ctx, "Badges fetched successfully", nil)
}

func (c *Controller) GetUserBadges(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetUserBadges endpoint called")
	return response.Success(ctx, "User badges fetched successfully", nil)
}
