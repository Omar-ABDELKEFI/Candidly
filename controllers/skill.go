package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

type SkillController struct{}

// CreateSkill godoc
// @Summary add new  skill
// @Description create new skill by json
// @Param skill body models.CreateSkillInput true "Add Skill"
// @Tags skill
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Skill
// @Router /skill [post]
func (h SkillController) CreateSkill(ctx *fiber.Ctx) error {
	var input models.CreateSkillInput
	validate := validator.New()
	log.Println("Hello from server")
	err := ctx.BodyParser(&input)
	validationError := validate.Struct(input)
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

	newSkill, err := services.CreateSkill(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   newSkill,
	})
}

// FindSkills godoc
// @Summary get skills
// @Description get all skills
// @Tags skill
// @Accept  json
// @Produce  json
// @Success 200 {array} models.SkillsResponse
// @Security Authorization
// @Router /skills [get]
func (h SkillController) FindSkills(ctx *fiber.Ctx) error {
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
