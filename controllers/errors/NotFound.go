package errors

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func NotFound(ctx *fiber.Ctx) error {
	CustomError := ReturnStruct{
		Status:  "Not Found",
		Time:    time.Now(),
		Message: "Item Not Found",
		Data:    "",
	}
	return ctx.Status(fiber.StatusNotFound).JSON(CustomError)
}
