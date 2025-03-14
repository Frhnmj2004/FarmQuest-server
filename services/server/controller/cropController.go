package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CropController struct {
	DB *gorm.DB
}

func NewCropController(db *gorm.DB) *CropController {
	return &CropController{DB: db}
}

func (cc *CropController) GetCrops(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetCrops endpoint called")

	// Get the category query parameter (default to "all")
	category := ctx.Query("category", "all")

	// Fetch crops based on category
	var crops []models.Crops
	query := cc.DB

	switch category {
	case "all":
		// No additional filtering; fetch all crops
		if err := query.Find(&crops).Error; err != nil {
			// utils.LogError("Failed to fetch all crops: %v", err)
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crops")
		}

	case "indoor", "outdoor":
		// Filter by Tags containing the category
		if err := query.Where("? = ANY(tags)", category).Find(&crops).Error; err != nil {
			// utils.LogError("Failed to fetch %s crops: %v", category, err)
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crops")
		}

	case "popular":
		// Fetch top 5 crops by Popularity, and optionally include crops tagged as "popular"
		if err := query.Where("popularity > 0 OR 'popular' = ANY(tags)").Order("popularity DESC").Limit(5).Find(&crops).Error; err != nil {
			// utils.LogError("Failed to fetch popular crops: %v", err)
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch popular crops")
		}

	default:
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid category parameter")
	}

	// Map to response type
	cropResponses := make([]types.CropResponse, len(crops))
	for i, crop := range crops {
		// Determine primary category for display (prioritize based on query)
		displayCategory := category
		if category == "all" {
			// Default to "indoor" or "outdoor" based on Tags
			for _, tag := range crop.Tags {
				if tag == "indoor" {
					displayCategory = "indoor"
					break
				} else if tag == "outdoor" {
					displayCategory = "outdoor"
					break
				}
			}
		}

		cropResponses[i] = types.CropResponse{
			ID:         crop.ID,
			Name:       crop.Name,
			Type:       crop.Type,
			Category:   displayCategory,
			ImageURL:   "https://example.com/images/" + crop.Name + ".jpg", // Dummy URL
			IsFavorite: false,                                              // Requires user context and favorites table
		}
	}

	return response.Success(ctx, "Crops retrieved successfully", cropResponses)
}

func (c *Controller) GetCropDetails(ctx *fiber.Ctx) error {
	return response.Success(ctx, "Crop details not implemented", nil)
}
