package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pebeasley/leaves-api/controllers/errors"
	"github.com/pebeasley/leaves-api/services"
)

type UserIDBody struct {
	ID string `validate:"required"`
}

func GetUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		user, err := services.GetUserById(id)

		if err != nil {
			return ctx.Status(fiber.StatusNotFound).JSON(err.Error())
		}
		if len(user.ID) < 1 {
			return errors.NotFound(ctx)
		}
		return ctx.Status(fiber.StatusOK).JSON(user)
	}
}

type NewUser struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
}

func UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}

func GetUserBudgets() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if len(id) == 0 {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Must include ID parameter",
				"data":    nil,
			})
		}

		foundBudgets, err := services.GetBudgetsByUserId(id)
		if err != nil {
			ctx.Status(fiber.StatusInternalServerError)
			return ctx.JSON(fiber.Map{
				"status":  "Failed",
				"message": err.Error(),
				"data":    nil,
			})
		}

		ctx.Status(fiber.StatusOK)
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Budgets found",
			"data":    foundBudgets,
		})
	}
}
