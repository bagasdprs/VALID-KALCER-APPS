package main

import (
	"log"
	"valid-kalcer-backend/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
		// 1. Load .env File
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}

		// 2. Connect to Database
		database.ConnectDB()

		// 3. Initialize Fiber App (Framework Web App)
		app := fiber.New(fiber.Config{
			AppName: "Valid Kalcer API v1",
		})

		// 4. Setup CORS
		app.Use(cors.New(cors.Config{
			AllowOrigins: "http://localhost:3000",
			AllowHeaders: "Origin, Content-Type, Accept",
			AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		}))

		// 5. Setup Routes
		app.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status": "Success",
				"message": "Welcome to Valid Kalcer API v1, your Server is ONLINE!",

			})
		})

		// 6. Start Server
		log.Println("Starting server on port 8080")
		log.Fatal(app.Listen(":8080"))
}