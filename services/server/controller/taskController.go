package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

func NewTaskController(db *gorm.DB) *TaskController {
	return &TaskController{DB: db}
}

func (tc *TaskController) GetTasks(ctx *fiber.Ctx) error {
	// utils.LogInfo("GetTasks endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Task listing not implemented")
}

func (tc *TaskController) CreateTask(ctx *fiber.Ctx) error {
	// utils.LogInfo("CreateTask endpoint called")
	return response.Error(ctx, fiber.StatusNotImplemented, "Task creation not implemented")
}

func (tc *TaskController) CompleteTask(ctx *fiber.Ctx) error {
	// utils.LogInfo("CompleteTask endpoint called, ID: %s", c.Params("id"))
	return response.Error(ctx, fiber.StatusNotImplemented, "Task completion not implemented")
}
