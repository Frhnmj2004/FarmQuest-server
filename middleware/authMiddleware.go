package middleware

// import (
// 	"os"
// 	"strings"

// 	// "github.com/Frhnmj2004/FarmQuest-server/utils"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/golang-jwt/jwt/v5"
// )

// // AuthMiddleware verifies JWT tokens and sets the user ID in the context
// func AuthMiddleware() fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		// Get the Authorization header
// 		authHeader := ctx.Get("Authorization")
// 		if authHeader == "" {
// 			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Authorization header required",
// 			})
// 		}

// 		// Check if header starts with "Bearer " and extract token
// 		const bearerPrefix = "Bearer "
// 		if !strings.HasPrefix(authHeader, bearerPrefix) {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid authorization format",
// 			})
// 		}
// 		tokenStr := strings.TrimPrefix(authHeader, bearerPrefix)

// 		// Parse and validate the JWT token
// 		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
// 			// Validate signing method
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
// 			}
// 			// Return the secret key for validation
// 			return []byte(os.Getenv("JWT_SECRET")), nil
// 		})

// 		if err != nil {
// 			utils.LogError("Failed to parse JWT: %v", err)
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 				"error": "Invalid or expired token",
// 			})
// 		}

// 		// Check if token is valid and extract claims
// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			// Extract user_id from claims and set it in context
// 			userID, ok := claims["user_id"].(float64) // JWT stores numbers as float64
// 			if !ok {
// 				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 					"error": "Invalid token claims",
// 				})
// 			}
// 			c.Locals("user_id", int(userID))
// 			return c.Next()
// 		}

// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": "Invalid token",
// 		})
// 	}
// }
