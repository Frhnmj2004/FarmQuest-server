package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type RewardController struct {
	DB *gorm.DB
}

func NewRewardController(db *gorm.DB) *RewardController {
	return &RewardController{DB: db}
}

func (rc *RewardController) GetRewards(c *fiber.Ctx) error {
	utils.LogInfo("GetRewards endpoint called")
	return c.JSON(fiber.Map{
		"message": "Reward listing not implemented",
	})
}

func (rc *RewardController) ClaimReward(c *fiber.Ctx) error {
	utils.LogInfo("ClaimReward endpoint called")
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Reward claiming not implemented",
	})
}
