package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/common"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
	"os"
	"strconv"
	"time"
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
	var testCandidates []models.TestCandidate
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
	testCandidate.TestID = id

	validate := validator.New()
	validationError := validate.Struct(testCandidate)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}
	testCandidates = append(testCandidates, testCandidate)
	// Create testCandidate
	emailsDuplicate, newTestCandidate, errDuplicate := services.CreateTestCandidate(testCandidates)
	if errDuplicate != nil {
		log.Println("Error ", errDuplicate)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":           errDuplicate,
			"emailsDuplicate": emailsDuplicate,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"test":   newTestCandidate,
	})
}

// CalculateScore godoc
// @Summary calculate a test score
// @Description calculate score by path  and update a status test
// @Param idTestCandidate path string true "idTestCandidate"
// @Tags test_candidate
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TestCandidate
// @Router /score/{idTestCandidate} [post]
func (h TestCandidateController) CalculateScore(ctx *fiber.Ctx) error {

	idTestCandidateEncrypted := ctx.Params("idTestCandidate")
	idOfTestCandidate := common.AesDecrypt(idTestCandidateEncrypted, os.Getenv("key"))
	log.Println("Hello from testCandidate")
	testId, errIdTest, candidateId, errIdCandidate := common.GetTestCandidate(idOfTestCandidate)
	if errIdTest != nil || errIdCandidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorIdTest":      errIdTest,
			"errorIdCandidate": errIdCandidate,
		})
	}
	testCandidate, err := services.CalculateScore(candidateId, testId)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   testCandidate,
	})
}

// FindTestsCandidates godoc
// @Summary get candidates and their tests
// @Description get candidates and their tests
// @Tags test_candidate
// @Accept  json
// @Produce  json
// @Success 200 {array} models.TestsCandidatesResponse
// @Security Authorization
// @Router /testscandidates [get]
func (h TestCandidateController) FindTestsCandidates(ctx *fiber.Ctx) error {
	testsCandidates, err := services.FindTestsCandidates()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   testsCandidates,
	})
}

// StartTest godoc
// @Summary get test information
// @Description get test information
// @id StartTest
// @Tags test_candidate
// @Param idTestCandidate path string true "idTestCandidate"
// @Produce  json
// @Success 200 {array} models.StartTest
// @Security Authorization
// @Router /startTest/{idTestCandidate} [get]
func StartTest(ctx *fiber.Ctx) error {
	idTestCandidateEncrypted := ctx.Params("idTestCandidate")
	idTestCandidate := common.AesDecrypt(idTestCandidateEncrypted, os.Getenv("key"))
	log.Println(idTestCandidate, "idTestcandidat")
	idTest, errIdTest, idCandidate, errIdCandidate := common.GetTestCandidate(idTestCandidate)
	if errIdTest != nil || errIdCandidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorIdTest":      errIdTest,
			"errorIdCandidate": errIdCandidate,
		})
	}
	log.Println(idTest, "idTestidTest")
	log.Println(idCandidate, "idCandidateidCandidate")
	testCandidate, err := services.StartTest(idTest, idCandidate)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	log.Println(testCandidate.CreatedAt, "testCandidate.CreatedAt")

	createdAt := testCandidate.CreatedAt
	log.Println(createdAt, "createdAt")

	log.Println(time.Duration(testCandidate.TimeLimit).Hours(), "testCandidate.TimeLimit")
	if time.Now().After(createdAt.Add(time.Hour * 24 * time.Duration(testCandidate.TimeLimit))) {
		var testStatus models.UpdateTestStatus

		testStatus.TestStatus = "canceled"
		testStatusOutput, _ := services.UpdateTestStatus(idTest, idCandidate, testStatus)
		if testStatusOutput.TestStatus == testStatus.TestStatus {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status":        "canceled",
				"testCandidate": testCandidate,
				"time":          createdAt.Add(time.Hour * 24 * time.Duration(testCandidate.TimeLimit)),
			})
		}

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   testCandidate,
	})

}

// FindTestCandidate godoc
// @Summary get tests candidates
// @Description get questions of a test
// @Tags test
// @Param idTestCandidate path string true "idTestCandidate"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Question
// @Security Authorization
// @Router /quiz/{idTestCandidate} [get]
func (h TestCandidateController) FindQuiz(ctx *fiber.Ctx) error {
	idTestCandidateEncrypted := ctx.Params("idTestCandidate")
	idTestCandidate := common.AesDecrypt(idTestCandidateEncrypted, os.Getenv("key"))
	testId, err, _, _ := common.GetTestCandidate(idTestCandidate)
	println(testId, "testIdDinfquizcontrollers")
	if err != nil {
		log.Println("could not find testID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorIdTest": err,
		})

	}
	quiz, err := services.FindQuiz(testId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   quiz,
	})

}

// Update Test Status godoc
// @Summary Update test status
// @Description Update test status
// @Accept  json
// @Produce  json
// @Param idTestCandidate path string true "idTestCandidate"
// @Param  testStatus body models.UpdateTestStatus true "test status"
// @Success 200 {object} models.UpdateTestStatusOutput
// @Router /quiz/status/{idTestCandidate} [Patch]
func (h TestCandidateController) UpdateTestStatus(ctx *fiber.Ctx) error {
	idTestCandidateEncrypted := ctx.Params("idTestCandidate")
	idTestCandidate := common.AesDecrypt(idTestCandidateEncrypted, os.Getenv("key"))
	testId, errIdTest, candidateId, errIdCandidate := common.GetTestCandidate(idTestCandidate)
	if errIdTest != nil || errIdCandidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorIdTest":      errIdTest,
			"errorIdCandidate": errIdCandidate,
		})
	}
	var testStatus models.UpdateTestStatus
	err := ctx.BodyParser(&testStatus)
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	testStatusOutput, err := services.UpdateTestStatus(testId, candidateId, testStatus)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   testStatusOutput,
	})

}

// Update Test Status godoc
// @Summary Update current question
// @Description Update current question
// @Accept  json
// @Produce  json
// @Param idTestCandidate path string true "idTestCandidate"
// @Param  currentQuestion body models.UpdateCurrentQuestion true "current question"
// @Success 200 {object} models.UpdateCurrentQuestionOutput
// @Router /quiz/currentQuestion/{idTestCandidate} [Patch]
func (h TestCandidateController) UpdateCurrentQuestion(ctx *fiber.Ctx) error {
	idTestCandidateEncrypted := ctx.Params("idTestCandidate")
	idTestCandidate := common.AesDecrypt(idTestCandidateEncrypted, os.Getenv("key"))
	testId, errIdTest, candidateId, errIdCandidate := common.GetTestCandidate(idTestCandidate)
	if errIdTest != nil || errIdCandidate != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorIdTest":      errIdTest,
			"errorIdCandidate": errIdCandidate,
		})
	}
	var currentQuestion models.UpdateCurrentQuestion
	err := ctx.BodyParser(&currentQuestion)
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	currentQuestionOutput, err := services.UpdateCurrentQuestion(testId, candidateId, currentQuestion)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "succes",
		"data":   currentQuestionOutput,
	})

}
