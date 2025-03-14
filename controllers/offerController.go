package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type OfferController struct {
	DB *gorm.DB
}

func NewOfferController(db *gorm.DB) *OfferController {
	return &OfferController{DB: db}
}

func (oc *OfferController) GetPersonalizedOffers(c *fiber.Ctx) error {
	utils.LogInfo("GetPersonalizedOffers endpoint called")
	return c.JSON(fiber.Map{
		"message": "Personalized offers not implemented",
	})
}
