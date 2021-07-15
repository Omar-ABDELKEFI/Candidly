package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

func CreateAnswer(ctx *fiber.Ctx) error {
	var answer models.Answer
	log.Println("Hello from server")
	err := ctx.BodyParser(&answer)
	validate := validator.New()
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	validationError := validate.Struct(answer)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}
	// Create Answer
	newAnswer, err := services.CreateAnswer(answer)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"answer": newAnswer,
	})
}
