package controller

import (
	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/response"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/types"
	"github.com/gofiber/fiber/v2"
)

func (c *Controller) GetQuestions(ctx *fiber.Ctx) error {
	return response.Success(ctx, "Questions retrieved successfully", []types.GetQuestionsResponse{})
}
