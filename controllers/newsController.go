package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type NewsController struct {
	DB *gorm.DB
}

func NewNewsController(db *gorm.DB) *NewsController {
	return &NewsController{DB: db}
}

func (nc *NewsController) GetNews(c *fiber.Ctx) error {
	utils.LogInfo("GetNews endpoint called")
	return c.JSON(fiber.Map{
		"message": "News listing not implemented",
	})
}

func (nc *NewsController) CreateNews(c *fiber.Ctx) error {
	utils.LogInfo("CreateNews endpoint called")
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "News creation not implemented",
	})
}
