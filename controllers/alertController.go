package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AlertController struct {
	DB *gorm.DB
}

func NewAlertController(db *gorm.DB) *AlertController {
	return &AlertController{DB: db}
}

func (ac *AlertController) GetAlerts(c *fiber.Ctx) error {
	utils.LogInfo("GetAlerts endpoint called")
	return c.JSON(fiber.Map{
		"message": "Alert listing not implemented",
	})
}

func (ac *AlertController) MarkAlertAsRead(c *fiber.Ctx) error {
	utils.LogInfo("MarkAlertAsRead endpoint called, ID: %s", c.Params("id"))
	return c.JSON(fiber.Map{
		"message": "Mark alert as read not implemented",
	})
}
