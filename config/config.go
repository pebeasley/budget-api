package config

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper

	errorHandler fiber.ErrorHandler
	fiber        *fiber.Config
}

var defaultErrorHandler = func(c *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Set error message
	message := err.Error()

	// Check if it's a fiber.Error type
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	// TODO: Check return type for the client, JSON, HTML, YAML or any other (API vs web)

	// Return HTTP response
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	c.Status(code)

	// Render default error view
	err = c.Render("errors/"+strconv.Itoa(code), fiber.Map{"message": message})
	if err != nil {
		return c.SendString(message)
	}
	return err
}

func New() *Config {
	config := &Config{
		Viper: viper.New(),
	}

	config.setDefaults()

	// Select the .env file
	config.SetConfigName(".env")
	config.SetConfigType("dotenv")
	config.AddConfigPath(".")

	config.AutomaticEnv()

	config.SetErrorHandler(defaultErrorHandler)

	config.setFiberConfig()
	return config
}

func (config *Config) SetErrorHandler(errorHandler fiber.ErrorHandler) {
	config.errorHandler = errorHandler
}

func (config *Config) setDefaults() {
	fmt.Println("Setting config defaults...")
}

func (config *Config) setFiberConfig() {
	config.fiber = &fiber.Config{}
}

func (config *Config) GetFiberConfig() *fiber.Config {
	return config.fiber
}
