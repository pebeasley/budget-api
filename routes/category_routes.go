package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
)

func setupCategoryRoutes(api fiber.Router) {
	categories := api.Group("/categories")
	categories.Get("/:id", Controller.GetCategory())
}
