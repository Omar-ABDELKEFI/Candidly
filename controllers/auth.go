package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/dev/test/models"
	"github.com/dev/test/services"
	"log"
)

type AuthController struct{}

// Login godoc
// @Summary Login to the app
// @Description Login to the app
// @Param user body models.LoginInput true "Login"
// @Tags Login
// @Accept  json
// @Produce  json
// @Success 200 {string} string "ok"
// @Router /login [post]
func (h AuthController) Login(ctx *fiber.Ctx) error {
	var user models.LoginInput
	log.Println("Hello from server")
	//validate := validator.New()
	err := ctx.BodyParser(&user)
	//validationError := validate.Struct(user)
	//element:=validationError
	if user.Password == "" && user.Email == "" {
		log.Println("Error : password and email required")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "password_and_email_required",
		})
	}
	if user.Password == "" {
		log.Println("Error : password required")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "password_required",
		})
	}
	if user.Email == "" {
		log.Println("Error : password required")
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "email_required",
		})
	}

	if err != nil {
		log.Println("Error : Invalid Json format")
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}
	token, err := services.ValidateLogin(user.Email, user.Password)
	if err != nil {
		log.Println("Error ", err.Error())
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid_login",
		})
	}

	fmt.Printf("%+v\n", ctx.Get("Authorization"))

	log.Println(ctx.Request().Header.Header()[0], "ctx.Request().Header.IsGet() in controllers/auth")
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"token":  token,
		"ctx":    ctx.Request().Header.IsGet(),
	})
}
