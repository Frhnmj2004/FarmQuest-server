package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {
	req := types.LoginRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	return response.Success(ctx, "Login successful", types.LoginResponse{})
}

func (c *Controller) Register(ctx *fiber.Ctx) error {
	req := types.RegisterRequest{}

	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	return response.Success(ctx, "Register successful", types.RegisterResponse{})
}
