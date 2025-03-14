package controller

import (
	"fmt"
	"github.com/Frhnmj2004/FarmQuest-server.git/models"
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"time"
)

// GetTasks retrieves tasks for the authenticated user
func (c *Controller) GetTasks(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	// Get task type from query parameter
	taskType := ctx.Query("type", "all")

	// Build query
	query := c.db.Where("user_id = ?", userID)
	if taskType != "all" {
		query = query.Where("type = ?", taskType)
	}

	var tasks []models.Task
	if err := query.Find(&tasks).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch tasks: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to fetch tasks")
	}

	return response.Success(ctx, "Tasks retrieved successfully", fiber.Map{
		"tasks": tasks,
	})
}

// CreateTask creates a new task for the authenticated user
func (c *Controller) CreateTask(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	// Parse request body
	var task models.Task
	if err := ctx.BodyParser(&task); err != nil {
		return response.Error(ctx, fiber.StatusBadRequest, "Invalid request body")
	}

	// Set user ID and timestamps
	task.UserID = uint(userID)
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	// Validate task fields
	if task.Title == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Task title is required")
	}
	if task.Type == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Task type is required")
	}
	if task.Points < 0 {
		return response.Error(ctx, fiber.StatusBadRequest, "Task points cannot be negative")
	}

	// Create task
	if err := c.db.Create(&task).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to create task: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to create task")
	}

	return response.Success(ctx, "Task created successfully", fiber.Map{
		"task": task,
	})
}

// CompleteTask marks a task as completed
func (c *Controller) CompleteTask(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("user_id").(int)
	if !ok {
		return response.Error(ctx, fiber.StatusUnauthorized, "User not authenticated")
	}

	taskID := ctx.Params("id")
	if taskID == "" {
		return response.Error(ctx, fiber.StatusBadRequest, "Task ID is required")
	}

	// Find task and verify ownership
	var task models.Task
	if err := c.db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to fetch task: %s", err.Error()))
		return response.Error(ctx, fiber.StatusNotFound, "Task not found")
	}

	if task.IsCompleted {
		return response.Error(ctx, fiber.StatusBadRequest, "Task is already completed")
	}

	// Update task status
	task.IsCompleted = true
	task.CompletedAt = time.Now()
	task.UpdatedAt = time.Now()

	if err := c.db.Save(&task).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to update task: %s", err.Error()))
		return response.Error(ctx, fiber.StatusInternalServerError, "Failed to complete task")
	}

	// Update user points
	if err := c.db.Model(&models.User{}).Where("id = ?", userID).
		UpdateColumn("points", gorm.Expr("points + ?", task.Points)).Error; err != nil {
		c.logger.Error(fmt.Errorf("Failed to update user points: %s", err.Error()))
		// Don't return error here, task is still completed
	}

	return response.Success(ctx, "Task completed successfully", fiber.Map{
		"task": task,
	})
}
