package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/muhammadfikri4/go-plate/config"
	"github.com/muhammadfikri4/go-plate/database"
	"github.com/muhammadfikri4/go-plate/middlewares"
	"github.com/muhammadfikri4/go-plate/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config.LoadConfig()
	database.ConnectDB()

	app := fiber.New()

	middlewares.SetupCORS(app)

	routes.SetupRoutesApp(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
