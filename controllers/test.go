package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

// CreateTest godoc
// @Summary add new Test
// @Description create new Test by json
// @Param Test body models.CreateTestInput true "Add Test"
// @Tags test
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Test
// @Router /my-tests [post]
func CreateTest(ctx *fiber.Ctx) error {
	var input models.CreateTestInput
	log.Println("Hello from server")
	err := ctx.BodyParser(&input)
	validate := validator.New()
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	validationError := validate.Struct(input)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}
	// Create Test
	newTest, err := services.CreateTest(input)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"test":   newTest,
	})
}

// FindTests godoc
// @Summary get tests
// @Description get tests by skill
// @Param type query uint64 false "tests search by type"
// @Tags test
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Test
// @Security Authorization
// @Router /tests [get]
func FindTests(ctx *fiber.Ctx) error {
	skillsId := ctx.Query("skills")
	tests, err := services.FindTests(skillsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   tests,
	})

}
