package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"time"
)

// GetHealth checks the health status of the server
func (c *Controller) GetHealth(ctx *fiber.Ctx) error {
	return response.Success(ctx, "Server is healthy", fiber.Map{
		"status": "ok",
		"time":   time.Now(),
	})
}
