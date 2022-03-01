package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
)

func setupAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", Controller.Login())
	auth.Post("/register", Controller.Register())
}
