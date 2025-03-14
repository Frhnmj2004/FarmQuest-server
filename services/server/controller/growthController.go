package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetGrowthStatus(ctx *fiber.Ctx) error {
	farmID := ctx.Params("id")
	if farmID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Farm ID is required")
	}

	return response.Success(ctx, "Growth status retrieved successfully", []types.GetGrowthStatusResponse{})
}
