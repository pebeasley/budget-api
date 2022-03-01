package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
)

func SetupRoutesV1(api fiber.Router) {
	setupHomeRoutes(api)
	setupAuthRoutes(api)
	setupUserRoutes(api)
	setupBudgetRoutes(api)
	setupTransactionRoutes(api)
	setupAccountRoutes(api)
	setupCategoryRoutes(api)
}

func setupHomeRoutes(api fiber.Router) {
	home := api.Group("/")
	home.Get("/", Controller.GetHomePage())
}
