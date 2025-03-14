package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CropController struct {
	DB *gorm.DB
}

func NewCropController(db *gorm.DB) *CropController {
	return &CropController{DB: db}
}

func (cc *CropController) GetCrops(c *fiber.Ctx) error {
	utils.LogInfo("GetCrops endpoint called")
	return c.JSON(fiber.Map{
		"message": "Crop listing not implemented",
	})
}

func (cc *CropController) GetCropDetails(c *fiber.Ctx) error {
	utils.LogInfo("GetCropDetails endpoint called, ID: %s", c.Params("id"))
	return c.JSON(fiber.Map{
		"message": "Crop details not implemented",
	})
}
