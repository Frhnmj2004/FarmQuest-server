package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"time"
)

// GetRewards retrieves all rewards for the authenticated user
func (c *Controller) GetRewards(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var rewards []models.Reward
	if err := c.db.Where("user_id = ?", userID).Find(&rewards).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch rewards: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch rewards")
	}

	return response.Success(ctx, "Rewards retrieved successfully", fiber.Map{
		"rewards": rewards,
	})
}

// GetReward retrieves a specific reward by ID
func (c *Controller) GetReward(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	rewardID := ctx.Params("id")
	if rewardID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Reward ID is required")
	}

	var reward models.Reward
	if err := c.db.Where("id = ? AND user_id = ?", rewardID, userID).First(&reward).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch reward: %s", err.Error()))
		return response.Error(ctx, fiber.StatusNotFound, "Reward not found")
	}

	return response.Success(ctx, "Reward retrieved successfully", fiber.Map{
		"reward": reward,
	})
}

// ClaimReward claims a reward for the authenticated user
func (c *Controller) ClaimReward(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	rewardID := ctx.Params("id")
	if rewardID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Reward ID is required")
	}

	var reward models.Reward
	if err := c.db.Where("id = ? AND user_id = ? AND claimed = false", rewardID, userID).First(&reward).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch reward: %s", err.Error()))
		return response.Error(ctx, fiber.StatusNotFound, "Reward not found or already claimed")
	}

	// Update reward status
	reward.Claimed = true
	reward.ClaimedAt = time.Now()
	reward.UpdatedAt = time.Now()

	if err := c.db.Save(&reward).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to update reward: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to claim reward")
	}

	return response.Success(ctx, "Reward claimed successfully", fiber.Map{
		"reward": reward,
	})
}
