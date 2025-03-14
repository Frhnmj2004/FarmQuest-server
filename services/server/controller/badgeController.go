package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// GetBadges retrieves all available badges
func (c *Controller) GetBadges(ctx *fiber.Ctx) error {
	var badges []models.Badge
	if err := c.db.Find(&badges).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch badges: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch badges")
	}

	return response.Success(ctx, "Badges fetched successfully", fiber.Map{
		"badges": badges,
	})
}

// GetUserBadges retrieves badges earned by the authenticated user
func (c *Controller) GetUserBadges(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var userBadges []models.UserBadge
	if err := c.db.Preload("Badge").Where("user_id = ?", userID).Find(&userBadges).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch user badges: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch user badges")
	}

	return response.Success(ctx, "User badges fetched successfully", fiber.Map{
		"badges": userBadges,
	})
}
