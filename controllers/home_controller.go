package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	validate *validator.Validate = validator.New()
)

func GetHomePage() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(&fiber.Map{"hello": "world"})
	}
}
