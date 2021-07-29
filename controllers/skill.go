package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

// CreateSkill godoc
// @Summary add new  skill
// @Description create new skill by json
// @Param skill body models.Skill true "Add Skill"
// @Tags skill
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Skill
// @Router /skill [post]
func CreateSkill(ctx *fiber.Ctx) error {
	var skill models.Skill
	validate := validator.New()
	log.Println("Hello from server")
	err := ctx.BodyParser(&skill)
	validationError := validate.Struct(skill)
	log.Println(validationError, "validationError")
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	// Create skill

	if err := services.CreateSkill(skill); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   skill,
	})
}

// FindSkills godoc
// @Summary get skills
// @Description get all skills
// @Tags skill
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Skill
// @Security Authorization
// @Router /skills [get]
func FindSkills(ctx *fiber.Ctx) error {
	skills, err := services.FindSkills()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   skills,
	})

}
