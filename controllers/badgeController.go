package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BadgeController struct {
	DB *gorm.DB
}

func NewBadgeController(db *gorm.DB) *BadgeController {
	return &BadgeController{DB: db}
}

func (bc *BadgeController) GetBadges(c *fiber.Ctx) error {
	utils.LogInfo("GetBadges endpoint called")
	return c.JSON(fiber.Map{
		"message": "Badge listing not implemented",
	})
}

func (bc *BadgeController) GetUserBadges(c *fiber.Ctx) error {
	utils.LogInfo("GetUserBadges endpoint called")
	return c.JSON(fiber.Map{
		"message": "User badge listing not implemented",
	})
}
