package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// GetAlerts retrieves alerts for the authenticated user
func (c *Controller) GetAlerts(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var alerts []models.Alert
	if err := c.db.Where("user_id = ?", userID).Find(&alerts).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch alerts: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch alerts")
	}

	return response.Success(ctx, "Alerts retrieved successfully", fiber.Map{
		"alerts": alerts,
	})
}

// MarkAlertAsRead marks an alert as read
func (c *Controller) MarkAlertAsRead(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	alertID := ctx.Params("id")
	if alertID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Alert ID is required")
	}

	result := c.db.Model(&models.Alert{}).Where("id = ? AND user_id = ?", alertID, userID).Update("is_read", true)
	if result.Error != nil {
		c.logger.Error(fmt.Errorf("Failed to mark alert as read: %s", result.Error.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to mark alert as read")
	}

	if result.RowsAffected == 0 {
		return response.Error(ctx, fiber.StatusNotFound, "Alert not found")
	}

	return response.Success(ctx, "Alert marked as read", nil)
}
