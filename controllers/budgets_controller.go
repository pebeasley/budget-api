package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pebeasley/leaves-api/services"
)

func CreateBudget() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Params("id"))
		return ctx.SendString("Not Implemented")
	}
}

func GetBudget() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if len(id) == 0 {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Bad Request",
				"data":    nil,
			})
		}
		return ctx.SendString("Not Implemented")
	}
}

func DeleteBudget() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if len(id) == 0 {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Bad Request",
				"data":    nil,
			})
		}
		err := services.DeleteBudget(id)
		if err != nil {
			ctx.Status(fiber.StatusNotFound)
			return ctx.JSON(fiber.Map{
				"status":  fiber.StatusNotFound,
				"message": "Not Found",
				"data":    nil,
			})
		}
		ctx.Status(fiber.StatusOK)
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "BudgetDeleted",
			"data":    nil,
		})
	}
}

func GetBudgetAccounts() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.SendString("Not Implemented")
	}
}

func GetBudgetCategories() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if len(id) == 0 {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"status":  "Failed",
				"message": "Bad Request",
				"data":    nil,
			})
		}
		categories, err := services.GetCategoriesByBudgetID(id)

		if err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(fiber.Map{
				"status":  "Failed",
				"message": "Bad Request",
				"data":    err.Error(),
			})
		}

		if len(categories) == 0 {
			ctx.Status(fiber.StatusNotFound)
			return ctx.JSON(fiber.Map{
				"status":  "Failed",
				"message": "No results found",
				"data":    categories,
			})
		}

		ctx.Status(fiber.StatusOK)
		return ctx.JSON(fiber.Map{
			"status":  "Success",
			"message": "Success",
			"data":    categories,
		})
	}
}
