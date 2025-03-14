package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FarmController struct {
	DB *gorm.DB
}

func NewFarmController(db *gorm.DB) *FarmController {
	return &FarmController{DB: db}
}

func (c *Controller) CreateFarm(ctx *fiber.Ctx) error {
	// utils.LogInfo("CreateFarm endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Farm creation not implemented")
}

func (c *Controller) GetFarms(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetFarms endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Farm listing not implemented")
}

func (c *Controller) GetFarm(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetFarm endpoint called, ID: %s", c.Params("id"))
	return response.Error(ctx, fiber.StatusNotImplemented, "Farm retrieval not implemented")
}

func (c *Controller) UpdateFarm(ctx *fiber.Ctx) error {
	// utils.LogInfo("UpdateFarm endpoint called, ID: %s", c.Params("id"))
	return response.Error(ctx, fiber.StatusNotImplemented, "Farm update not implemented")
}
