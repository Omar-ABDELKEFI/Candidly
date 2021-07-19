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
	app.Post("/my-tests", controllers.CreateTest)
	app.Get("/tests", controllers.FindTests)
	app.Post("/candidate", controllers.CreateCandidate)
	app.Post("/my-tests/:id/questions", controllers.CreateTestQuestion)
	app.Post("/score/:id", controllers.CalculateScore)
	app.Post("/answers", controllers.CreateAnswer)

	app.Post("/my-tests/candidates/:id", controllers.CreateTestCandidate)
	questionController := new(controllers.QuestionController)
	question := app.Group("/question")
	{
		question.Post("/edit", questionController.CreateQuestion)
		question.Get("/", questionController.FindQuestion)
	}

}
