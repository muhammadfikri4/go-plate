package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfikri4/go-plate/app/routes"
	"github.com/muhammadfikri4/go-plate/utils"
)

func SetupRoutesApp(app *fiber.App) {
	routes.UserRoute(app.Group("/users"))

	// Default route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Welcome To REST API 🚀"})
	})

	// 404 Route
	app.Use(func(c *fiber.Ctx) error {
		responseHandler := &utils.ResponseHandler{}

		return responseHandler.NotFound(c, []string{"Route not found, check again your endpoint"})
	})
}
