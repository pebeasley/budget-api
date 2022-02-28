package routes

import (
	"github.com/gofiber/fiber/v2"
	Controller "github.com/pebeasley/leaves-api/controllers"
	"github.com/pebeasley/leaves-api/middleware"
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

func setupAuthRoutes(api fiber.Router) {
	auth := api.Group("/auth")
	auth.Post("/login", Controller.Login())
	auth.Post("/register", Controller.Register())
}

func setupUserRoutes(api fiber.Router) {
	users := api.Group("/users")
	users.Use(middleware.Protected())
	users.Get("/:id", Controller.GetUser())
	users.Get("/:id/budgets", Controller.GetUserBudgets())
	users.Patch("/:id", Controller.UpdateUser())
}

func setupBudgetRoutes(api fiber.Router) {
	budgets := api.Group("/budgets")
	budgets.Post("/", Controller.CreateBudget())
	budgets.Get("/:id/", Controller.GetBudget())
	budgets.Get("/:id/accounts", middleware.Protected(), Controller.GetBudgetAccounts())
	budgets.Get("/:id/categories", middleware.Protected(), Controller.GetBudgetCategories())
	budgets.Delete("/:id", Controller.DeleteBudget())
}

func setupTransactionRoutes(api fiber.Router) {
	transactions := api.Group("/transactions")
	transactions.Get("/", Controller.GetTransactions())
	transactions.Post("/", Controller.AddTransaction())
	transactions.Get("/:id", Controller.GetTransaction())
	transactions.Delete("/:id", Controller.DeleteTransaction())
	transactions.Patch("/:id", Controller.UpdateTransaction())
}

func setupAccountRoutes(api fiber.Router) {
	accounts := api.Group("/accounts")
	accounts.Get("/:id", Controller.GetAccount())
	accounts.Post("/:id", Controller.UpdateAccount())
	accounts.Delete("/:id", Controller.DeleteAccount())
	accounts.Get("/:id/transactions", Controller.GetAccountTransactions())
}

func setupCategoryRoutes(api fiber.Router) {
	categories := api.Group("/categories")
	categories.Get("/:id", Controller.GetCategory())
}
