package controllers

import "github.com/gofiber/fiber/v2"

func GetHomePage() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(&fiber.Map{"hello": "world"})
	}
}
