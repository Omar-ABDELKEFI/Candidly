package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
	"strconv"
)

// CreateTestQuestion godoc
// @Summary add a question to test
// @Description add a question to test by json
// @Param test_id path int true "test id"
// @Param test_question body models.TestQuestion true "Add question to test"
// @Tags question_test
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TestQuestion
// @Router /my-tests/:id/questions [post]
func CreateTestQuestion(ctx *fiber.Ctx) error {
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
	testQuestion.TestId = id
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
		"test":   newTestQuestion,
	})
}
