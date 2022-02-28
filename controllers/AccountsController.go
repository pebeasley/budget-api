package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pebeasley/leaves-api/services"
)

func GetAccount() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}

func UpdateAccount() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}

func DeleteAccount() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}

func GetAccountTransactions() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		account := ctx.Params("id")
		tx, _ := services.GetTransactionsByAccount(account)
		return ctx.Status(fiber.StatusOK).JSON(tx)
	}
}
