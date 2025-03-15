package controller

import (
	"errors"

	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/utils"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {
	req := types.LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Find user by email
	var user models.User
	if err := c.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process login")
		}
	}

	// Verify password
	if err := utils.VerifyPassword(user.Password, req.Password); err != nil {
		return response.Error(ctx, fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(int(user.ID))
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process login")
	}

	return response.Success(ctx, "Login successful", types.LoginResponse{
		Token: token,
	})
}

func (c *Controller) Register(ctx *fiber.Ctx) error {
	req := types.RegisterRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Check if user exists
	var user models.User
	if err := c.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process registration")
		}
	} else {
		return response.Error(ctx, fiber.StatusConflict, "User already exists")
	}

	if err := c.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process registration")
		}
	} else {
		return response.Error(ctx, fiber.StatusConflict, "User already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process registration")
	}

	user = models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := c.db.Create(&user).Error; err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process registration")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(int(user.ID))
	if err != nil {
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process registration")
	}

	return response.Success(ctx, "Register successful", types.RegisterResponse{
		Token: token,
	})
}
