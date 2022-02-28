package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetTransactions() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).
			JSON(fiber.Map{
				"message": "Not Implemented",
			})
	}
}

func AddTransaction() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}

func GetTransaction() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Not Implemented"})
	}
}

func UpdateTransaction() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}

func DeleteTransaction() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}
