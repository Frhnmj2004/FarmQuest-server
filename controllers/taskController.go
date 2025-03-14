package controllers

import (
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

func NewTaskController(db *gorm.DB) *TaskController {
	return &TaskController{DB: db}
}

func (tc *TaskController) GetTasks(c *fiber.Ctx) error {
	utils.LogInfo("GetTasks endpoint called")
	return c.JSON(fiber.Map{
		"message": "Task listing not implemented",
	})
}

func (tc *TaskController) CreateTask(c *fiber.Ctx) error {
	utils.LogInfo("CreateTask endpoint called")
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Task creation not implemented",
	})
}

func (tc *TaskController) CompleteTask(c *fiber.Ctx) error {
	utils.LogInfo("CompleteTask endpoint called, ID: %s", c.Params("id"))
	return c.JSON(fiber.Map{
		"message": "Task completion not implemented",
	})
}
