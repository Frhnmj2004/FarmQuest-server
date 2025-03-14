package response

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(fiber.Map{"message": message, "data": data, "error": false})
}

func Error(c *fiber.Ctx, status int, errorMessage string) error {
	return c.Status(status).JSON(fiber.Map{"message": errorMessage, "data": nil, "error": true})
}
