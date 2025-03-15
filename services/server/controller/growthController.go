package controller

import (
	"time"

	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetGrowthStatus(ctx *fiber.Ctx) error {
	farmID := ctx.Params("farm_id")
	if farmID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm ID is required")
	}

	// Dummy data for growth stages
	return response.Success(ctx, "Growth status retrieved successfully", types.GetGrowthStatusResponse{
		Stages: []struct {
			Description string    `json:"description"`
			Status      string    `json:"status"`
			Date        time.Time `json:"date,omitempty"`
		}{
			{
				Description: "The plant has been planted and is in the early stages of growth.",
				Status:      "planted",
				Date:        time.Now(),
			},
			{
				Description: "Plant is growing well with healthy leaves.",
				Status:      "growing",
				Date:        time.Now().Add(7 * 24 * time.Hour),
			},
			{
				Description: "Plant is ready for harvesting.",
				Status:      "harvesting",
				Date:        time.Now().Add(14 * 24 * time.Hour),
			},
		},
	})
}
