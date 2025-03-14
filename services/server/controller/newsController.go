package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type NewsController struct {
	DB *gorm.DB
}

func NewNewsController(db *gorm.DB) *NewsController {
	return &NewsController{DB: db}
}

func (c *Controller) GetNews(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetNews endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "News listing not implemented")
}

func (c *Controller) CreateNews(ctx *fiber.Ctx) error {
	// utils.LogInfo("CreateNews endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "News creation not implemented")
}
