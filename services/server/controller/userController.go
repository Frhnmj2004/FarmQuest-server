package controller

import (
	"fmt"

	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/utils"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register handles user registration
func (c *Controller) Register(ctx *fiber.Ctx) error {

	req := types.RegisterRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.logger.Error(fmt.Errorf("Failed to hash password: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process password")
	}

	// Create user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user", // Default role
		Points:   0,      // Initial points
		Balance:  0,      // Initial balance
	}
	if err := c.db.Create(&user).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to create user: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to register user")
	}

	return response.Success(ctx, "User registered successfully", fiber.Map{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"role":       user.Role,
		"points":     user.Points,
		"balance":    user.Balance,
		"created_at": user.CreatedAt,
	})
}

func (c *Controller) Login(ctx *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Find user by email
	var user models.User
	if err := c.db.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Error(ctx, fiber.StatusUnauthorized, "Invalid email or password")
		}
		c.logger.Error(fmt.Errorf("Failed to find user: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to process login")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return response.Error(ctx, fiber.StatusUnauthorized, "Invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(int(user.ID))
	if err != nil {
		c.logger.Error(fmt.Errorf("Failed to generate token: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to generate token")
	}

	return response.Success(ctx, "Login successful", fiber.Map{
		"token": token,
		"user": fiber.Map{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"role":       user.Role,
			"points":     user.Points,
			"balance":    user.Balance,
			"created_at": user.CreatedAt,
		},
	})
}

func (c *Controller) GetProfile(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var user models.User
	if err := c.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Error(ctx, fiber.StatusNotFound, "User not found")
		}
		c.logger.Error(fmt.Errorf("Failed to fetch user: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch profile")
	}

	var profile models.Profile
	if err := c.db.Where("user_id = ?", userID).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			profile = models.Profile{UserID: uint(userID)} // Create empty profile if not found
		} else {
			c.logger.Error(fmt.Errorf("Failed to fetch profile: %s", err.Error()))
			return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch profile")
		}
	}

	return response.Success(ctx, "Profile retrieved successfully", fiber.Map{
		"full_name":  profile.FullName,
		"address":    profile.Address,
		"phone":      profile.Phone,
		"avatar_url": profile.AvatarURL,
		"bio":        profile.Bio,
	})
}

func (c *Controller) UpdateProfile(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var req struct {
		FullName  string `json:"full_name"`
		Address   string `json:"address"`
		Phone     string `json:"phone"`
		AvatarURL string `json:"avatar_url"`
		Bio       string `json:"bio"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	var profile models.Profile
	result := c.db.Where("user_id = ?", userID).FirstOrCreate(&profile, models.Profile{UserID: uint(userID)})
	if result.Error != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch or create profile: %s", result.Error.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update profile")
	}

	// Update fields if provided
	if req.FullName != "" {
		profile.FullName = req.FullName
	}
	if req.Address != "" {
		profile.Address = req.Address
	}
	if req.Phone != "" {
		profile.Phone = req.Phone
	}
	if req.AvatarURL != "" {
		profile.AvatarURL = req.AvatarURL
	}
	if req.Bio != "" {
		profile.Bio = req.Bio
	}

	if err := c.db.Save(&profile).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to update profile: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update profile")
	}

	return response.Success(ctx, "Profile updated successfully", fiber.Map{
		"full_name":  profile.FullName,
		"address":    profile.Address,
		"phone":      profile.Phone,
		"avatar_url": profile.AvatarURL,
		"bio":        profile.Bio,
	})
}
