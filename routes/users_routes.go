package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
	"github.com/pebeasley/leaves-api/middleware"
)

func setupUserRoutes(api fiber.Router) {
	users := api.Group("/users")
	users.Use(middleware.Protected())
	users.Get("/:id", Controller.GetUser())
	users.Get("/:id/budgets", Controller.GetUserBudgets())
	users.Patch("/:id", Controller.UpdateUser())
}
