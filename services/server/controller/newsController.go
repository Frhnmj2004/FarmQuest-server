package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"time"
)

// GetNews retrieves all news articles
func (c *Controller) GetNews(ctx *fiber.Ctx) error {
	var news []models.News
	if err := c.db.Find(&news).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch news: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch news")
	}

	return response.Success(ctx, "News retrieved successfully", fiber.Map{
		"news": news,
	})
}

// CreateNews creates a new news article (admin only)
func (c *Controller) CreateNews(ctx *fiber.Ctx) error {
	// Check admin status
	isAdmin, ok := ctx.Locals("is_admin").(bool)
	if !ok || !isAdmin {
		return response.Error(ctx, fiber.StatusUnauthorized, "Admin access required")
	}

	var article models.News
	if err := ctx.BodyParser(&article); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	if err := c.db.Create(&article).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to create article: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to create article")
	}

	return response.Success(ctx, "Article created successfully", fiber.Map{
		"article": article,
	})
}
