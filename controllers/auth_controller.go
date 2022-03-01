package controllers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pebeasley/leaves-api/controllers/errors"
	"github.com/pebeasley/leaves-api/services"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

type LoginBodyStruct struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

func Login() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		fmt.Println("Hit Auth Controller Login Func")
		LoginBody := new(LoginBodyStruct)
		if err := ctx.BodyParser(LoginBody); err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(
				fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"data":    nil,
				},
			)
		}

		if err := validate.Struct(LoginBody); err != nil {
			ctx.Status(fiber.StatusBadRequest)
			return ctx.JSON(
				fiber.Map{
					"status":  fiber.StatusBadRequest,
					"message": err.Error(),
					"data":    nil,
				})
		}

		user, err := services.GetUserByEmail(LoginBody.Email)

		if err != nil || len(user.FirstName) == 0 {
			return errors.NotFound(ctx)
		}

		if !CheckPasswordHash(LoginBody.Password, user.Password) {
			ctx.Status(fiber.StatusUnauthorized)
			return ctx.JSON(
				fiber.Map{
					"status":  fiber.StatusUnauthorized,
					"message": "Invalid username or password",
					"data":    nil,
				})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = user.Email
		claims["user_id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("SECRET"))
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}

		ctx.Status(fiber.StatusOK)
		return ctx.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Sign in successful",
			"data": fiber.Map{
				"token":    t,
				"id":       user.ID,
				"username": user.Email,
			},
		})
	}
}

func Register() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Not Implemented")
	}
}
