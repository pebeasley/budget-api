package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
	"github.com/pebeasley/leaves-api/middleware"
)

func setupTransactionRoutes(api fiber.Router) {
	transactions := api.Group("/transactions")
	transactions.Use(middleware.Protected())
	transactions.Get("/", Controller.GetTransactions())
	transactions.Post("/", Controller.AddTransaction())
	transactions.Get("/:id", Controller.GetTransaction())
	transactions.Delete("/:id", Controller.DeleteTransaction())
	transactions.Patch("/:id", Controller.UpdateTransaction())
}
