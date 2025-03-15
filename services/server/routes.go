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
		authRoutes := mainRoutes.Group("/auth")
		{
			authRoutes.Post("/login", baseController.Login)
			authRoutes.Post("/register", baseController.Register)
		}

		mainRoutes.Get("/health", baseController.GetHealth)
		mainRoutes.Get("/crops", baseController.GetCrops)
		mainRoutes.Get("/crops/dropdown", baseController.GetCropsDropdown)
		mainRoutes.Get("/crops/:id", baseController.GetCrop)
		mainRoutes.Get("/farms", baseController.GetFarms)
		mainRoutes.Get("/farms/:id", baseController.GetFarm)
		mainRoutes.Get("/farms/:id/growth", baseController.GetGrowthStatus)
		mainRoutes.Get("/questions", baseController.GetQuestions)
	}

	return r
}
