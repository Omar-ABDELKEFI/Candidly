package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/dev/test/models"
	"github.com/dev/test/services"
	"log"
	"strconv"
)

type TestQuestionController struct{}

// CreateTestQuestion godoc
// @Summary add a question to test
// @Description add a question to test by json
// @Param id path string true "id"
// @Param test_question body models.TestQuestion true "Add question to test"
// @Tags question_test
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TestQuestion
// @Router /my-tests/{id}/questions [post]
func (h TestQuestionController) CreateTestQuestion(ctx *fiber.Ctx) error {
	var testQuestion models.TestQuestion
	log.Println("Hello from server")
	err := ctx.BodyParser(&testQuestion)
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	testQuestion.TestID = id
	validate := validator.New()
	validationError := validate.Struct(testQuestion)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}

	// Create testQuestion
	newTestQuestion, err := services.CreateTestQuestion(testQuestion)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   newTestQuestion,
	})
}

// DeleteTestQuestion godoc
// @Summary delete a question from test
// @Description delete a question from test by json
// @Param test_question body models.TestQuestionDelete true "test_question"
// @Tags question_test
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /my-tests/questions/delete [Delete]
func (h TestQuestionController) DeleteTestQuestion(ctx *fiber.Ctx) error {
	var input models.TestQuestionDelete
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
	err = services.DeleteTestQuestion(input.TestID, input.QuestionID)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
	})
}
