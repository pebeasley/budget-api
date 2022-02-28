package errors

import "github.com/gofiber/fiber/v2"

func NotImplemented() fiber.Map {
	return fiber.Map{"message": "Not Implemented"}
}
