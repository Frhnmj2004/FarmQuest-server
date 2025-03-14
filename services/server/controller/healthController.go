package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) HealthCheck(fctx *fiber.Ctx) error {
	return response.Success(fctx, "Server is running healthy", nil)
}
