package main

import (
	"log"
	"os"

	"github.com/Frhnmj2004/FarmQuest-server/config"
	"github.com/Frhnmj2004/FarmQuest-server/routes"
	"github.com/Frhnmj2004/FarmQuest-server/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database connection
	db := config.InitDB()
	if db == nil {
		log.Fatal("Failed to initialize database")
	}

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			utils.LogError("Unhandled error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		},
	})

	// Middleware
	app.Use(cors.New())    // Enable CORS
	app.Use(logger.New())  // Log HTTP requests
	app.Use(recover.New()) // Recover from panics

	// Setup routes
	routes.SetupRoutes(app)

	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	utils.LogInfo("Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		utils.LogError("Server failed to start: %v", err)
		os.Exit(1)
	}
}
