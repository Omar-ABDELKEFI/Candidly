package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
	"strconv"
)

type TestController struct{}

// CreateTest godoc
// @Summary add new Test
// @Description create new Test by json
// @Param Test body models.TestRequest true "Add Test"
// @Tags test
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TestResponse
// @Router /my-tests [post]
func (h TestController) CreateTest(ctx *fiber.Ctx) error {
	var input models.Test
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
// @Success 200 {array} models.TestResponse
// @Security Authorization
// @Router /tests [get]
func (h TestController) FindTests(ctx *fiber.Ctx) error {
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

// UpdateTest godoc
// @Summary update Test
// @id updateTest
// @Description update Test by json and path
// @Param Test body models.TestRequest true "Update Test"
// @Param test_id path int true "Update Test"
// @Tags test
// @Accept  json
// @Produce  json
// @Success 200 {object} models.TestResponse
// @Router /my-tests/{test_id} [post]
func (h TestController) UpdateTest(ctx *fiber.Ctx) error {
	var input models.Test
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
	testId, err := strconv.ParseUint(ctx.Params("idTest"), 10, 64)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	newTest, err := services.UpdateTest(input, testId)
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

// getMyTests godoc
// @Summary update Test
// @id getMyTests
// @Description get my-tests
// @Tags test
// @Produce  json
// @Success 200 {object} models.MyTests
// @Router /my-tests/getTest [get]
func (h TestController) GetMyTests(ctx *fiber.Ctx) error {

	log.Println("Hello from server")

	newTest, err := services.GetMyTests()
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

// getTest godoc
// @Summary get test by id
// @id getTest
// @Description get test by id
// @Param id path string true "test id"
// @Tags test
// @Produce  json
// @Success 200 {object} models.MyTests
// @Router /my-tests/{id} [get]
func (h TestController) GetTest(ctx *fiber.Ctx) error {
	testId, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Println("could not find testID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorIdTest": err,
		})
	}

	test, err := services.GetTest(testId)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   test,
	})
}

// clone test godoc
// @Summary clone a test
// @id cloneTest
// @Description clone test
// @Param id path string true "test id"
// @Param expectedTime body models.CloneTestInput true "expected Time"
// @Tags test
// @Produce  json
// @Success 200 {object} models.MyTests
// @Router /my-tests/clone/{id} [post]
func (h TestController) CloneTest(ctx *fiber.Ctx) error {
	testId, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		log.Println("could not find testID")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errorIdTest": err,
		})
	}
	var input models.CloneTestInput
	err = ctx.BodyParser(&input)
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
	clonedTest, err := services.CloneTest(testId, input)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "SUCCESS",
		"data":   clonedTest,
	})
}
