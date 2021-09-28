package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/common"
	_ "github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
	"strings"
)

type QuestionController struct{}

// CreateQuestion godoc
// @Summary add new  question
// @Description create new question by json
// @Param question body models.CreateQuestionInput true "Add question"
// @Tags question
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Question
// @Router /questions/edit [post]
func (h QuestionController) CreateQuestion(ctx *fiber.Ctx) error {
	var input models.CreateQuestionInput
	log.Println("Hello from server")
	err := ctx.BodyParser(&input)
	errors, validationError := common.ValidateStruct(&input)
	log.Println(errors, "validationError")
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	// Create question
	if input.SkillName != "" {
		skillId, err := services.FindOrCreateSkill(input.SkillName)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err,
			})
		}
		input.SkillID = uint64(skillId)
	}
	question, err := services.CreateQuestion(input)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   question,
	})
}

// FindQuestion godoc
// @Summary find a question
// @Description find a question by type or difficulty
// @Param type query string false "question search by type"
// @Param difficulty query string false "question search by difficulty"
// @Tags question
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Question
// @Router /questions [get]
func (h QuestionController) FindQuestion(ctx *fiber.Ctx) error {
	var tabSort []string
	var tabDifficulty []string
	log.Println(tabSort, "tabSort in controllers/question")
	log.Println(len(tabSort), "(len(tabSort) in controllers/question ")
	if sort := ctx.Query("types"); len(sort) != 0 {
		tabSort = strings.Split(sort, ",")

	}
	if difficulty := ctx.Query("difficulties"); len(difficulty) != 0 {
		tabDifficulty = strings.Split(difficulty, ",")
	}

	questions, err := services.FindQuestion(tabSort, tabDifficulty)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   questions,
	})
}
