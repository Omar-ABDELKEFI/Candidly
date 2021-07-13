package server

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/tekab-dev/tekab-test/controllers"
	"github.com/tekab-dev/tekab-test/docs"
)

func Router(app *fiber.App) {
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	app.Get("/swagger/*", swagger.Handler)
	app.Post("/login", controllers.Login)
	app.Post("/questions", controllers.CreateQuestion)
	//app.Get("/question/:type?")
	app.Post("/skill", controllers.CreateSkill)
	app.Post("/test", controllers.CreateTest)
	app.Post("/candidat", controllers.CreateCandidat)

}
