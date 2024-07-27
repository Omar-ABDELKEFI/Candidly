package controllers

import (
	_ "github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/dev/test/models"
	"github.com/dev/test/services"
	"log"
)

type CandidateController struct{}

// CreateCandidate godoc
// @Summary add new  Candidate
// @Description create new Candidate by json
// @Param candidate body []models.CandidateRequest true "candidate data"
// @Tags Candidate
// @Accept  json
// @Produce  json
// @Success 200 {object} models.CandidateResponse
// @Router /candidate [post]
func (h CandidateController) CreateCandidate(ctx *fiber.Ctx) error {
	var candidate []models.Candidate
	log.Println("server is running")
	err := ctx.BodyParser(&candidate)
	//validate := validator.New()
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	// todo create Candidate validation
	//validationError := validate.Struct(candidate)
	//if validationError != nil {
	//	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"error": validationError.Error(),
	//	})
	//}
	// Create candidate
	newTestCandidate, newCandidate, emailsDuplicate, errDuplicate, err := services.CreateCandidate(candidate)
	if len(errDuplicate) != 0 {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":            err,
			"newCandidate":     newCandidate,
			"errDuplicate":     errDuplicate,
			"emailsDuplicate":  emailsDuplicate,
			"newTestCandidate": newTestCandidate,
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":           "SUCCESS",
		"candidate":        newCandidate,
		"newTestCandidate": newTestCandidate,
	})
}
