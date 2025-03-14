package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetRewards(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetRewards endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Reward listing not implemented")
}

func (c *Controller) ClaimReward(ctx *fiber.Ctx) error {
	// utils.LogInfo("ClaimReward endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Reward claiming not implemented")
}
