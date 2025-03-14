package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

func NewOfferController(db *gorm.DB) *Controller {
	return &Controller{db: db}
}

func (c *Controller) GetPersonalizedOffers(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetPersonalizedOffers endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Personalized offers not implemented")
}

func (c *Controller) GetOffers(ctx *fiber.Ctx) error {
	var offers []models.Offer
	if err := c.db.Find(&offers).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch offers: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch offers")
	}

	return response.Success(ctx, "Offers retrieved successfully", fiber.Map{
		"offers": offers,
	})
}

func (c *Controller) GetOffer(ctx *fiber.Ctx) error {
	offerID := ctx.Params("id")
	if offerID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Offer ID is required")
	}

	var offer models.Offer
	if err := c.db.First(&offer, offerID).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch offer: %s", err.Error()))
		return response.Error(ctx, fiber.StatusNotFound, "Offer not found")
	}

	return response.Success(ctx, "Offer retrieved successfully", fiber.Map{
		"offer": offer,
	})
}

func (c *Controller) CreateOffer(ctx *fiber.Ctx) error {
	// Check admin status
	isAdmin, ok := ctx.Locals("is_admin").(bool)
	if !ok || !isAdmin {
		return response.Error(ctx, fiber.StatusUnauthorized, "Admin access required")
	}

	var offer models.Offer
	if err := ctx.BodyParser(&offer); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	offer.CreatedAt = time.Now()
	offer.UpdatedAt = time.Now()

	if err := c.db.Create(&offer).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to create offer: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to create offer")
	}

	return response.Success(ctx, "Offer created successfully", fiber.Map{
		"offer": offer,
	})
}
