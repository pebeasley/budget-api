package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	configuration "github.com/pebeasley/leaves-api/config"
	"github.com/pebeasley/leaves-api/database"
	"github.com/pebeasley/leaves-api/routes"
)

type App struct {
	*fiber.App

	DB *database.Database
}

func main() {
	config := configuration.New()
	app := App{
		App: fiber.New(*config.GetFiberConfig()),
	}

	// Initialize database
	err := database.New(&database.DatabaseConfig{
		Driver:   config.GetString("DB_DRIVER"),
		Host:     config.GetString("DB_HOST"),
		Username: config.GetString("DB_USERNAME"),
		Password: config.GetString("DB_PASSWORD"),
		Port:     config.GetInt("DB_PORT"),
		Database: config.GetString("DB_DATABASE"),
	})

	if err != nil {
		fmt.Println("Failed to connect to database")
	}

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	routes.SetupRoutesV1(app.Group("/api"))

	listenError := app.Listen(":9080")

	if listenError != nil {
		fmt.Println(listenError)
		fmt.Println(err.Error())
	}
}
