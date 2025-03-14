package response

import "github.com/gofiber/fiber/v2"

func Success(ctx *fiber.Ctx, message string, data interface{}) error {
	return ctx.JSON(fiber.Map{"message": message, "data": data, "error": false})
}

func Error(ctx *fiber.Ctx, status int, errorMessage string) error {
	return ctx.Status(status).JSON(fiber.Map{"message": errorMessage, "data": nil, "error": true})
}
