package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
	"strconv"
)

type TestCandidateController struct{}

// CreateTestCandidate godoc
// @Summary add new  test_candidate
// @Description create test_candidate by json and path
// @Param test_id path int true "Add candidate"
// @Param test_candidate body models.TestCandidate true "Add candidate"
// @Tags test_candidate
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TestCandidate
// @Router /my-tests/candidates/:test_id [post]
func (h TestCandidateController) CreateTestCandidate(ctx *fiber.Ctx) error {
	var testCandidate models.TestCandidate
	log.Println("Hello from testCandidate")
	err := ctx.BodyParser(&testCandidate)
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
	testCandidate.TestId = id
	validate := validator.New()
	validationError := validate.Struct(testCandidate)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}

	// Create testCandidate
	newTestCandidat, err := services.CreateTestCandidate(testCandidate)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"test":   newTestCandidat,
	})
}

// CalculateScore godoc
// @Summary calculate a test score
// @Description calculate score by path  and update a status test
// @Param test_candidate_id path  int true "calculate score"
// @Tags test_candidate
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TestCandidate
// @Router /score [post]
func (h TestCandidateController) CalculateScore(ctx *fiber.Ctx) error {
	var testCandidate models.TestCandidate
	log.Println("Hello from testCandidate")
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	testCandidate, err = services.CalculateScore(id)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"test":   testCandidate,
	})
}
