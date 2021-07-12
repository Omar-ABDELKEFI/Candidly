package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

func CreateTest(ctx *fiber.Ctx) error {
	var test models.Test
	log.Println("Hello from server")
	validate := validator.New()
	err := ctx.BodyParser(&test)
	validationError := validate.Struct(test)
	log.Println("validationError", validationError)
	if err != nil || validationError != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	newTest, err := services.CreateTest(test)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"test":   newTest,
	})
}
