package routes

import (
	"github.com/Frhnmj2004/FarmQuest-server/controllers" // Adjust to your actual package path
	"github.com/Frhnmj2004/FarmQuest-server/middleware"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// SetupRoutes configures the API endpoints for the FarmQuest application
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Initialize controllers with database connection
	userCtrl := controllers.NewUserController(db)
	farmCtrl := controllers.NewFarmController(db)
	orderCtrl := controllers.NewOrderController(db)
	rewardCtrl := controllers.NewRewardController(db)
	newsCtrl := controllers.NewNewsController(db)
	alertCtrl := controllers.NewAlertController(db)
	badgeCtrl := controllers.NewBadgeController(db)
	taskCtrl := controllers.NewTaskController(db)
	offerCtrl := controllers.NewOfferController(db)
	cropCtrl := controllers.NewCropController(db)

	// Public routes (no authentication required)
	app.Post("/register", userCtrl.Register)
	app.Post("/login", userCtrl.Login)

	// Protected routes (require JWT authentication)
	api := app.Group("/api", middleware.AuthMiddleware())
	{
		// Crop routes
		api.Get("/crops", cropCtrl.GetCrops)           // List all crops with offers
		api.Get("/crops/:id", cropCtrl.GetCropDetails) // Details of a specific crop

		// Offer routes
		api.Get("/crops/offers", offerCtrl.GetPersonalizedOffers) // Personalized offers

		// User routes
		api.Get("/profile", userCtrl.GetProfile)
		api.Put("/profile", userCtrl.UpdateProfile)

		// Farm routes
		api.Post("/farms", farmCtrl.CreateFarm)
		api.Get("/farms", farmCtrl.GetFarms)
		api.Get("/farms/:id", farmCtrl.GetFarm)
		api.Put("/farms/:id", farmCtrl.UpdateFarm)

		// Order routes
		api.Post("/orders", orderCtrl.CreateOrder)
		api.Get("/orders", orderCtrl.GetOrders)
		api.Get("/orders/:id", orderCtrl.GetOrder)
		api.Put("/orders/:id/status", orderCtrl.UpdateOrderStatus)

		// Reward routes
		api.Get("/rewards", rewardCtrl.GetRewards)
		api.Post("/rewards/claim", rewardCtrl.ClaimReward)

		// News routes
		api.Get("/news", newsCtrl.GetNews)
		api.Post("/news", newsCtrl.CreateNews) // Admin only, add middleware if needed

		// Alert routes
		api.Get("/alerts", alertCtrl.GetAlerts)
		api.Put("/alerts/:id/read", alertCtrl.MarkAlertAsRead)

		// Badge routes
		api.Get("/badges", badgeCtrl.GetBadges)
		api.Get("/user-badges", badgeCtrl.GetUserBadges)

		// Task routes
		api.Get("/tasks", taskCtrl.GetTasks)
		api.Post("/tasks", taskCtrl.CreateTask)
		api.Put("/tasks/:id/complete", taskCtrl.CompleteTask)
	}
}
