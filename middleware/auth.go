package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

// Protected protect routes
func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte("SECRET"),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusForbidden)
		return c.JSON(fiber.Map{
			"status":  "error",
			"message": "Missing or malformed JWT",
			"data":    nil,
		})

	}

	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"status":  "error",
		"message": "Invalid or expired JWT",
		"data":    nil,
	})
}
