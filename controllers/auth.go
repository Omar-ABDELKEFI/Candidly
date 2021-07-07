package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
)

func Login(ctx *fiber.Ctx) error {
	var body models.User
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return nil
	}
	if body.Username != "loukil" || body.Password != "pass123" {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		return nil
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": body,
	})
}
