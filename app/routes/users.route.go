package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfikri4/go-plate/app/controllers"
	"github.com/muhammadfikri4/go-plate/app/dto"
	"github.com/muhammadfikri4/go-plate/app/repositories"
	"github.com/muhammadfikri4/go-plate/app/services"
	"github.com/muhammadfikri4/go-plate/middlewares"
)

func UserRoute(route fiber.Router) {
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	route.Get("/", userController.GetUsers)
	route.Get("/:userId", userController.GetUser)
	route.Post("/", middlewares.ValidateRequest(&dto.CreateUserDTO{}), userController.CreateUser)
	route.Patch(
		"/:userId",
		middlewares.ValidateRequest(&dto.UpdateUserDTO{}),
		userController.UpdateUser,
	)
	route.Delete("/:id", userController.DeleteUser)
}
