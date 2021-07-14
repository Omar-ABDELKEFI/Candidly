package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
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
// @Param question body models.Question true "Add question"
// @Tags question
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Question
// @Security Authorization
// @Router /questions/edit [post]
func (h QuestionController) CreateQuestion(ctx *fiber.Ctx) error {
	var question models.Question
	validate := validator.New()
	log.Println("Hello from server")
	err := ctx.BodyParser(&question)
	validationError := validate.Struct(question)
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

	if err := services.CreateQuestion(question); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
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
// @Security Authorization
// @Router /questions [get]
func (h QuestionController) FindQuestion(ctx *fiber.Ctx) error {
	var tabSort []string
	var tabDifficulty []string
	log.Println(tabSort, "tabSorttabSort")
	log.Println(len(tabSort), "(len(tabSort)(len(tabSort)")
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
		"status": "succes",
		"data":   questions,
	})
}
