package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// CreateFarm creates a new farm for the authenticated user
func (c *Controller) CreateFarm(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)
	
	var req types.CreateFarmRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Validate request fields
	if req.Name == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm name is required")
	}
	if len(req.Name) > 255 {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm name cannot exceed 255 characters")
	}
	if req.Description == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm description is required")
	}
	if len(req.Description) > 1000 {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm description cannot exceed 1000 characters")
	}
	if req.Location == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm location is required")
	}
	if len(req.Location) > 255 {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm location cannot exceed 255 characters")
	}
	if req.Area <= 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm area must be greater than 0")
	}

	// Validate crop exists
	var crop models.Crop
	if err := c.db.First(&crop, req.CropID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusNotFound, "Crop not found")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to validate crop")
	}

	now := time.Now()

	// Create farm
	farm := models.Farms{
		Name:        req.Name,
		UserID:      userID,
		CropID:      int(req.CropID),
		Description: req.Description,
		Location:    req.Location,
		Area:        req.Area,
		Status:      "planted",
		ImageURL:    crop.CroppedImageURL,
		PlantedAt:   now,
		Health:      100,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := c.db.Create(&farm).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to create farm")
	}

	// Return farm details
	return response.Success(ctx, "Farm created successfully", types.GetSimpleFarmResponse{
		ID:       farm.ID,
		Name:     farm.Name,
		Image:    farm.ImageURL,
		Status:   farm.Status,
		Location: farm.Location,
		Area:     farm.Area,
	})
}

// GetFarms retrieves all farms for the authenticated user
func (c *Controller) GetFarms(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)

	var farms []models.Farms
	if err := c.db.Where("user_id = ?", userID).Find(&farms).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch farms")
	}

	farmResponses := []types.GetSimpleFarmResponse{}
	for _, farm := range farms {
		farmResponses = append(farmResponses, types.GetSimpleFarmResponse{
			ID:       farm.ID,
			Name:     farm.Name,
			Image:    farm.ImageURL,
			Status:   farm.Status,
			Location: farm.Location,
			Area:     farm.Area,
		})
	}

	return response.Success(ctx, "Farms retrieved successfully", farmResponses)
}

// GetFarm retrieves a specific farm by ID
func (c *Controller) GetFarm(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)
	farmID := ctx.Params("id")
	if farmID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm ID is required")
	}

	var farm models.Farms
	if err := c.db.Where("id = ? AND user_id = ?", farmID, userID).First(&farm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusNotFound, "Farm not found")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch farm")
	}

	// Get related crop information
	var crop models.Crop
	if err := c.db.First(&crop, farm.CropID).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crop details")
	}

	// Calculate growth metrics
	growthDescription := "Just planted"
	if !farm.GrowingAt.IsZero() {
		growthDescription = "Growing steadily"
	}
	if !farm.HarvestAt.IsZero() {
		growthDescription = "Ready for harvest"
	}

	return response.Success(ctx, "Farm retrieved successfully", types.GetFarmResponse{
		ID:          farm.ID,
		Name:        farm.Name,
		Image:       crop.FullImageURL,
		Status:      farm.Status,
		Description: farm.Description,
		Location:    farm.Location,
		Health:      farm.Health,
		Area:        farm.Area,
		PlantedAt:   farm.PlantedAt,
		GrowingAt:   farm.GrowingAt,
		HarvestAt:   farm.HarvestAt,
		GrowthStatus: struct {
			Description string `json:"description"`
			Image       string `json:"image_url"`
		}{
			Description: growthDescription,
			Image:       crop.FullImageURL,
		},
		RelatedNews: []struct {
			Title       string `json:"title"`
			Description string `json:"description"`
			Image       string `json:"image_url"`
			Link        string `json:"link"`
		}{
			{
				Title:       "Tips for Growing " + crop.Name,
				Description: "Learn the best practices for growing " + crop.Name,
				Image:       crop.CroppedImageURL,
				Link:        fmt.Sprintf("/crops/%d", crop.ID),
			},
		},
	})
}

// UpdateFarm updates a farm's status and other fields
func (c *Controller) UpdateFarm(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)
	farmID := ctx.Params("id")
	if farmID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm ID is required")
	}

	var farm models.Farms
	if err := c.db.Where("id = ? AND user_id = ?", farmID, userID).First(&farm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusNotFound, "Farm not found")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch farm")
	}

	var req types.UpdateFarmRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	now := time.Now()

	// Update fields if provided
	if req.Name != "" {
		if len(req.Name) > 255 {
			return response.Error(ctx, fiber.StatusBadRequest, "Farm name cannot exceed 255 characters")
		}
		farm.Name = req.Name
	}
	if req.Description != "" {
		if len(req.Description) > 1000 {
			return response.Error(ctx, fiber.StatusBadRequest, "Farm description cannot exceed 1000 characters")
		}
		farm.Description = req.Description
	}
	if req.Location != "" {
		if len(req.Location) > 255 {
			return response.Error(ctx, fiber.StatusBadRequest, "Farm location cannot exceed 255 characters")
		}
		farm.Location = req.Location
	}
	if req.Status != "" {
		// Validate status transition
		validStatuses := map[string]bool{
			"planted":    true,
			"growing":    true,
			"harvesting": true,
			"completed":  true,
		}
		if !validStatuses[req.Status] {
			return response.Error(ctx, fiber.StatusBadRequest, "Invalid status")
		}

		// Update timestamps based on status transition
		switch req.Status {
		case "growing":
			if farm.Status == "planted" {
				farm.GrowingAt = now
			}
		case "harvesting":
			if farm.Status == "growing" {
				farm.HarvestAt = now
			}
		}
		farm.Status = req.Status
	}

	// Update health and area if provided
	if req.Health > 0 {
		if req.Health > 100 {
			return response.Error(ctx, fiber.StatusBadRequest, "Farm health cannot exceed 100")
		}
		farm.Health = req.Health
	}
	if req.Area > 0 {
		farm.Area = req.Area
	}

	farm.UpdatedAt = now
	if err := c.db.Save(&farm).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update farm")
	}

	return response.Success(ctx, "Farm updated successfully", types.GetSimpleFarmResponse{
		ID:       farm.ID,
		Name:     farm.Name,
		Image:    farm.ImageURL,
		Status:   farm.Status,
		Location: farm.Location,
		Area:     farm.Area,
	})
}
