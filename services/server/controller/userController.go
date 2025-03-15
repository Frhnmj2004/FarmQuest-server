package controller

import (
	"errors"

	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetProfile retrieves the user's profile information
func (c *Controller) GetProfile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)

	var user models.User
	if err := c.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusNotFound, "User not found")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch profile")
	}

	profile := types.UserProfile{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		AvatarURL: user.AvatarURL,
		Role:      user.Role,
		Points:    user.Points,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}

	return response.Success(ctx, "Profile retrieved successfully", profile)
}

// UpdateProfile updates the user's profile information
func (c *Controller) UpdateProfile(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)
	req := types.UpdateProfileRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	var user models.User
	if err := c.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusNotFound, "User not found")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch user")
	}

	// Check username uniqueness if changed
	if req.Username != "" && req.Username != user.Username {
		var existingUser models.User
		if err := c.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
			return response.Error(ctx, fiber.StatusConflict, "Username already taken")
		}
	}

	// Update fields if provided
	if req.Username != "" {
		user.Username = req.Username
	}
	if req.FullName != "" {
		user.FullName = req.FullName
	}
	if req.AvatarURL != "" {
		user.AvatarURL = req.AvatarURL
	}

	if err := c.db.Save(&user).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update profile")
	}

	return response.Success(ctx, "Profile updated successfully", nil)
}

// GetUserPoints retrieves the user's current points
func (c *Controller) GetUserPoints(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)

	var user models.User
	if err := c.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusNotFound, "User not found")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch user points")
	}

	return response.Success(ctx, "Points retrieved successfully", fiber.Map{
		"points": user.Points,
	})
}

// UpdateUserPoints updates the user's points
func (c *Controller) UpdateUserPoints(ctx *fiber.Ctx) error {
	userID := ctx.Locals("userID").(int)
	points := struct {
		Points int `json:"points"`
	}{}

	if err := ctx.BodyParser(&points); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	var user models.User
	if err := c.db.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusNotFound, "User not found")
		}
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch user")
	}

	user.Points = points.Points
	if err := c.db.Save(&user).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update points")
	}

	return response.Success(ctx, "Points updated successfully", fiber.Map{
		"points": user.Points,
	})
}
