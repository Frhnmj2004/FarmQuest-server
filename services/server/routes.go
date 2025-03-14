package server

import (
	"context"

	"github.com/Frhnmj2004/FarmQuest-server.git/pkg/logger"
	"github.com/Frhnmj2004/FarmQuest-server.git/services/server/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func SetupRoutes(_ context.Context, db *gorm.DB, logger logger.Logger) *fiber.App {
	r := fiber.New()

	r.Use(cors.New())

	baseController := controller.NewBaseController(db, logger)
	mainRoutes := r.Group("/api")
	{
		mainRoutes.Get("/health", baseController.HealthCheck)
	}

	return r
}
