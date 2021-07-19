package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

// CreateCandidate godoc
// @Summary add new  candidate
// @Description create new candidate by json
// @Param candidate body models.Candidate true "Add candidate"
// @Tags candidate
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Candidate
// @Router /candidate [post]
func CreateCandidate(ctx *fiber.Ctx) error {
	var candidate models.Candidate
	log.Println("server is running")
	err := ctx.BodyParser(&candidate)
	validate := validator.New()
	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	validationError := validate.Struct(candidate)
	if validationError != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": validationError.Error(),
		})
	}
	// Create candidate
	candidate, err = services.CreateCandidate(candidate)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":   "SUCCESS",
		"candidat": candidate,
	})
}
