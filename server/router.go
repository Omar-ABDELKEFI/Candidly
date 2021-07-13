package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/controllers"
)

func Router(app *fiber.App) {

	app.Post("/login", controllers.Login)
	app.Post("/question", controllers.CreateQuestion)
	app.Post("/skill", controllers.CreateSkill)
	app.Post("/test", controllers.CreateTest)
	app.Post("/candidat", controllers.CreateCandidat)

}
