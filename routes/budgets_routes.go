package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
	"github.com/pebeasley/leaves-api/middleware"
)

func setupBudgetRoutes(api fiber.Router) {
	budgets := api.Group("/budgets")
	budgets.Use(middleware.Protected())
	budgets.Post("/", Controller.CreateBudget())
	budgets.Get("/:id/", Controller.GetBudget())
	budgets.Get("/:id/accounts", Controller.GetBudgetAccounts())
	budgets.Get("/:id/categories", Controller.GetBudgetCategories())
	budgets.Delete("/:id", Controller.DeleteBudget())
}
