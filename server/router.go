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

	//app.Get("/question/:type?")
	app.Post("/skill", controllers.CreateSkill)
	app.Post("/test", controllers.CreateTest) // todo path:/my-tests
	app.Get("/tests", controllers.FindTests)
	app.Post("/candidat", controllers.CreateCandidat)
	app.Post("/my-tests/:id/questions", controllers.CreateTestQuestion)
	questionController := new(controllers.QuestionController)
	question := app.Group("/question")
	{
		question.Post("/edit", questionController.CreateQuestion)
		question.Get("/", questionController.FindQuestion)
	}

}
