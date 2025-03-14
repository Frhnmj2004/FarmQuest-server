package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FarmController struct {
	DB *gorm.DB
}

func NewFarmController(db *gorm.DB) *FarmController {
	return &FarmController{DB: db}
}

func (fc *FarmController) CreateFarm(c *fiber.Ctx) error {
	utils.LogInfo("CreateFarm endpoint called")
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Farm creation not implemented",
	})
}

func (fc *FarmController) GetFarms(c *fiber.Ctx) error {
	utils.LogInfo("GetFarms endpoint called")
	return c.JSON(fiber.Map{
		"message": "Farm listing not implemented",
	})
}

func (fc *FarmController) GetFarm(c *fiber.Ctx) error {
	utils.LogInfo("GetFarm endpoint called, ID: %s", c.Params("id"))
	return c.JSON(fiber.Map{
		"message": "Farm retrieval not implemented",
	})
}

func (fc *FarmController) UpdateFarm(c *fiber.Ctx) error {
	utils.LogInfo("UpdateFarm endpoint called, ID: %s", c.Params("id"))
	return c.JSON(fiber.Map{
		"message": "Farm update not implemented",
	})
}
