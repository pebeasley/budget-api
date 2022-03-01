package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
	"github.com/pebeasley/leaves-api/middleware"
)

func setupAccountRoutes(api fiber.Router) {
	accounts := api.Group("/accounts")
	accounts.Use(middleware.Protected())
	accounts.Get("/:id", Controller.GetAccount())
	accounts.Post("/:id", Controller.UpdateAccount())
	accounts.Delete("/:id", Controller.DeleteAccount())
	accounts.Get("/:id/transactions", Controller.GetAccountTransactions())
}
