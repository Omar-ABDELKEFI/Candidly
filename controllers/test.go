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
	err := ctx.BodyParser(&test)
	validate := validator.New()
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	validationError := validate.Struct(test)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}

	// Create Test
	newTest, err := services.CreateTest(test)
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
// @Tags tests
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
