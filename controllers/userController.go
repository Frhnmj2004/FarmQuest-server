package controllers

import (
	"github.com/gofiber/fiber"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	utils.LogInfo("Register endpoint called")
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Register not implemented",
	})
}

func (uc *UserController) Login(c *fiber.Ctx) error {
	utils.LogInfo("Login endpoint called")
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Login not implemented",
	})
}

func (uc *UserController) GetProfile(c *fiber.Ctx) error {
	utils.LogInfo("GetProfile endpoint called")
	return c.JSON(fiber.Map{
		"message": "Profile retrieval not implemented",
	})
}

func (uc *UserController) UpdateProfile(c *fiber.Ctx) error {
	utils.LogInfo("UpdateProfile endpoint called")
	return c.JSON(fiber.Map{
		"message": "Profile update not implemented",
	})
}
