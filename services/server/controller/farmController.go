package controller

import (
	"fmt"
	"time"

	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
)

// CreateFarm creates a new farm for the authenticated user
func (c *Controller) CreateFarm(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var farm models.Farms
	if err := ctx.BodyParser(&farm); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	farm.UserID = userID
	farm.Status = "planted"
	farm.PlantedAt = time.Now()

	if err := c.db.Create(&farm).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to create farm: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to create farm")
	}

	return response.Success(ctx, "Farm created successfully", fiber.Map{
		"farm": farm,
	})
}

// GetFarms retrieves all farms for the authenticated user
func (c *Controller) GetFarms(ctx *fiber.Ctx) error {
	// userID, ok := ctx.Locals("user_id").(int)
	// if !ok {
	// 	return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	// }

	// var farms []models.Farms
	// if err := c.db.Where("user_id = ?", userID).Find(&farms).Error; err != nil {
	// 	c.logger.Error(fmt.Errorf("Failed to fetch farms: %s", err.Error()))
	// 	return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch farms")
	// }

	return response.Success(ctx, "Farms retrieved successfully", []types.GetSimpleFarmResponse{})
}

// GetFarm retrieves a specific farm by ID
func (c *Controller) GetFarm(ctx *fiber.Ctx) error {
	// userID, ok := ctx.Locals("user_id").(int)
	// if !ok {
	// 	return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	// }

	farmID := ctx.Params("id")
	if farmID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm ID is required")
	}

	// var farm models.Farms
	// if err := c.db.Where("id = ? AND user_id = ?", farmID, userID).First(&farm).Error; err != nil {
	// 	c.logger.Error(fmt.Errorf("Failed to fetch farm: %s", err.Error()))
	// 	return response.Error(ctx, fiber.StatusNotFound, "Farm not found")
	// }

	return response.Success(ctx, "Farm retrieved successfully", types.GetFarmResponse{})
}

// UpdateFarm updates a farm's status
func (c *Controller) UpdateFarm(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	farmID := ctx.Params("id")
	if farmID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm ID is required")
	}

	var farm models.Farms
	if err := c.db.Where("id = ? AND user_id = ?", farmID, userID).First(&farm).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch farm: %s", err.Error()))
		return response.Error(ctx, fiber.StatusNotFound, "Farm not found")
	}

	// Parse update request
	var req struct {
		Status string `json:"status"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Update status and related fields
	farm.Status = req.Status
	if req.Status == "harvested" {
		farm.HarvestedAt = time.Now()
	}

	if err := c.db.Save(&farm).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to update farm: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update farm")
	}

	return response.Success(ctx, "Farm updated successfully", fiber.Map{
		"farm": farm,
	})
}
