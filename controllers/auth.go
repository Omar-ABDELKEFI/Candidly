package controllers

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/services"
	"log"
)

// Login godoc
// @Summary Login
// @Description login using email and password
// @Param user body models.User true "login"
// @Tags question
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Security Authorization
// @Router /login [post]

func Login(ctx *fiber.Ctx) error {
	var user models.User
	log.Println("Hello from server")
	validate := validator.New()
	err := ctx.BodyParser(&user)
	validationError := validate.Struct(user)
	if err != nil || validationError != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	token, err := services.ValidateLogin(user.Email, user.Password)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  token,
	})
}
