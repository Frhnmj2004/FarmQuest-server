package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetAlerts(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetAlerts endpoint called")
	return response.Success(ctx, "Alert listing not implemented", nil)
}

func (c *Controller) MarkAlertAsRead(ctx *fiber.Ctx) error {
	// utils.LogInfo("MarkAlertAsRead endpoint called, ID: %s", ctx.Params("id"))
	return response.Success(ctx, "Mark alert as read not implemented", nil)
}
