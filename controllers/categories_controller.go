package controllers

import "github.com/gofiber/fiber/v2"

func GetCategory() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}
