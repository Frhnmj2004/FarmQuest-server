package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OfferController struct {
	DB *gorm.DB
}

func NewOfferController(db *gorm.DB) *OfferController {
	return &OfferController{DB: db}
}

func (c *Controller) GetPersonalizedOffers(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetPersonalizedOffers endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Personalized offers not implemented")
}
