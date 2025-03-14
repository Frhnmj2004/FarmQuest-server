package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
)

// GetCrops retrieves crops based on category
func (c *Controller) GetCrops(ctx *fiber.Ctx) error {
	// Get the category query parameter (default to "all")
	category := ctx.Query("category", "all")

	// Fetch crops based on category
	var crops []models.Crop
	query := c.db

	switch category {
	case "all":
		// No additional filtering; fetch all crops
		if err := query.Find(&crops).Error; err != nil {
			c.logger.Error(fmt.Errorf("Failed to fetch all crops: %s", err.Error()))
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crops")
		}

	case "indoor", "outdoor":
		// Filter by category
		if err := query.Where("category = ?", category).Find(&crops).Error; err != nil {
			c.logger.Error(fmt.Errorf("Failed to fetch %s crops: %s", category, err.Error()))
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crops")
		}

	case "popular":
		// Fetch top 5 crops by popularity
		if err := query.Where("is_popular = ? OR popularity > ?", true, 0).Order("popularity DESC").Limit(5).Find(&crops).Error; err != nil {
			c.logger.Error(fmt.Errorf("Failed to fetch popular crops: %s", err.Error()))
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch popular crops")
		}

	default:
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid category parameter")
	}

	// Check if user is authenticated to get favorites
	userID, ok := ctx.Locals("user_id").(int)
	var favorites []models.CropFavorite
	if ok {
		if err := c.db.Where("user_id = ?", userID).Find(&favorites).Error; err != nil {
			c.logger.Error(fmt.Errorf("Failed to fetch user favorites: %s", err.Error()))
		}
	}

	// Create a map of favorite crop IDs for quick lookup
	favoriteCrops := make(map[uint]bool)
	for _, fav := range favorites {
		favoriteCrops[fav.CropID] = true
	}

	// Map to response
	var cropResponses []fiber.Map
	for _, crop := range crops {
		cropResponses = append(cropResponses, fiber.Map{
			"id":          crop.ID,
			"name":        crop.Name,
			"type":        crop.Type,
			"category":    crop.Category,
			"image_url":   crop.ImageURL,
			"is_favorite": favoriteCrops[crop.ID],
		})
	}

	return response.Success(ctx, "Crops retrieved successfully", fiber.Map{
		"crops": cropResponses,
	})
}

// GetCropDetails retrieves detailed information about a specific crop
func (c *Controller) GetCropDetails(ctx *fiber.Ctx) error {
	cropID := ctx.Params("id")
	if cropID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Crop ID is required")
	}

	var crop models.Crop
	if err := c.db.First(&crop, cropID).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch crop details: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crop details")
	}

	// Check if user is authenticated to get favorite status
	userID, ok := ctx.Locals("user_id").(int)
	isFavorite := false
	if ok {
		var favorite models.CropFavorite
		if err := c.db.Where("user_id = ? AND crop_id = ?", userID, cropID).First(&favorite).Error; err == nil {
			isFavorite = true
		}
	}

	return response.Success(ctx, "Crop details retrieved successfully", fiber.Map{
		"id":            crop.ID,
		"name":          crop.Name,
		"type":          crop.Type,
		"category":      crop.Category,
		"image_url":     crop.ImageURL,
		"rating":        crop.Rating,
		"review_count":  crop.ReviewCount,
		"description":   crop.Description,
		"water_need":    crop.WaterNeed,
		"sunlight_need": crop.SunlightNeed,
		"price":         crop.Price,
		"is_favorite":   isFavorite,
	})
}
