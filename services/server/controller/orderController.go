package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"time"
)

// GetOrders retrieves all orders for the authenticated user
func (c *Controller) GetOrders(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var orders []models.Order
	if err := c.db.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch orders: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch orders")
	}

	return response.Success(ctx, "Orders retrieved successfully", fiber.Map{
		"orders": orders,
	})
}

// GetOrder retrieves a specific order by ID
func (c *Controller) GetOrder(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	orderID := ctx.Params("id")
	if orderID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Order ID is required")
	}

	var order models.Order
	if err := c.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch order: %s", err.Error()))
		return response.Error(ctx, fiber.StatusNotFound, "Order not found")
	}

	return response.Success(ctx, "Order retrieved successfully", fiber.Map{
		"order": order,
	})
}

// CreateOrder creates a new order for the authenticated user
func (c *Controller) CreateOrder(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	var order models.Order
	if err := ctx.BodyParser(&order); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	order.UserID = uint(userID)
	order.Status = "pending"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	if err := c.db.Create(&order).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to create order: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to create order")
	}

	return response.Success(ctx, "Order created successfully", fiber.Map{
		"order": order,
	})
}

// UpdateOrderStatus updates the status of an order
func (c *Controller) UpdateOrderStatus(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	orderID := ctx.Params("id")
	if orderID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Order ID is required")
	}

	var order models.Order
	if err := c.db.Where("id = ? AND user_id = ?", orderID, userID).First(&order).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch order: %s", err.Error()))
		return response.Error(ctx, fiber.StatusNotFound, "Order not found")
	}

	// Parse update request
	var req struct {
		Status string `json:"status"`
	}
	if err := ctx.BodyParser(&req); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Update status
	order.Status = req.Status
	order.UpdatedAt = time.Now()

	if err := c.db.Save(&order).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to update order: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to update order")
	}

	return response.Success(ctx, "Order status updated successfully", fiber.Map{
		"order": order,
	})
}
