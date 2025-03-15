package controller

import (
	"strings"

	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
)

// GetCropsDropdown retrieves a simplified list of crops for dropdown menus
func (c *Controller) GetCropsDropdown(ctx *fiber.Ctx) error {
	var crops []models.Crop
	if err := c.db.Find(&crops).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crops")
	}

	// Format response for dropdown
	cropResponses := []struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Image string `json:"image"`
	}{}

	for _, crop := range crops {
		cropResponses = append(cropResponses, struct {
			ID    uint   `json:"id"`
			Name  string `json:"name"`
			Image string `json:"image"`
		}{
			ID:    crop.ID,
			Name:  crop.Name,
			Image: crop.CroppedImageURL,
		})
	}

	return response.Success(ctx, "Crops retrieved successfully", cropResponses)
}

// GetCrops retrieves crops based on category
func (c *Controller) GetCrops(ctx *fiber.Ctx) error {
	// Get the category query parameter (default to "all")
	tags := ctx.Query("tags", "all")

	var crops []models.Crop
	var err error

	if tags == "all" {
		// If no specific tags, fetch all crops
		err = c.db.Find(&crops).Error
	} else {
		// Split the tags string into an array
		tagArray := strings.Split(tags, ",")
		// Query crops that have any of the specified tags using ANY operator
		err = c.db.Where("tags && ?", pq.Array(tagArray)).Find(&crops).Error
	}

	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crops")
	}

	cropResponses := []types.GetSimpleCropResponse{}
	for _, crop := range crops {
		cropResponses = append(cropResponses, types.GetSimpleCropResponse{
			ID:              crop.ID,
			Name:            crop.Name,
			CroppedImageURL: crop.CroppedImageURL,
			FullImageURL:    crop.FullImageURL,
		})
	}

	return response.Success(ctx, "Crops fetched successfully", cropResponses)
}

// GetCrop retrieves detailed information about a specific crop
func (c *Controller) GetCrop(ctx *fiber.Ctx) error {
	cropID := ctx.Params("id")
	if cropID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Crop ID is required")
	}

	var crop models.Crop
	if err := c.db.First(&crop, cropID).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch crop details")
	}

	return response.Success(ctx, "Crop details retrieved successfully", types.GetCropResponse{
		ID:           crop.ID,
		Name:         crop.Name,
		FullImageURL: crop.FullImageURL,
		Description:  crop.Description,
		BasicNeeds:   crop.BasicNeeds,
		Tags:         crop.Tags,
		Price:        int64(crop.Price),
		Rating:       int64(crop.Rating),
	})
}
