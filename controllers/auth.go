package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

func Login(ctx *fiber.Ctx) error {
	var user models.User
	log.Println("hello from server")
	validate := validator.New()
	err := ctx.BodyParser(&user)

	validationError := validate.Struct(user)
	if err != nil || validationError != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		log.Println("Invalid Json format")
		return nil
	}
	if !services.ValidateLogin(user.Email, user.Password) {
		ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad Credentials",
		})
		log.Println("Invalid Credentials")

		return nil
	}
	token, err := services.CreateToken(user.Email)
	if err != nil {
		ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error": err,
		})
		log.Println("Invalid Token")
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"token":  token,
	})
}
