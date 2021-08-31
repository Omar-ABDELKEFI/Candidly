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
	//app.Get("/question/:type?")

	skillController := new(controllers.SkillController)
	questionController := new(controllers.QuestionController)
	testController := new(controllers.TestController)
	authController := new(controllers.AuthController)
	answerController := new(controllers.AnswerController)
	candidateController := new(controllers.CandidateController)
	testCandidateController := new(controllers.TestCandidateController)
	testQuestionController := new(controllers.TestQuestionController)
	score := app.Group("/score")
	{
		score.Post("/:idTestCandidate", testCandidateController.CalculateScore)

	}

	auth := app.Group("/login")
	{
		auth.Post("/", authController.Login)
	}
	answers := app.Group("/answers/:idTestCandidate")
	{
		answers.Post("/", answerController.CreateAnswer)
	}
	myTest := app.Group("/my-tests")

	{
		myTest.Get("/getTest", testController.GetMyTests)
		myTest.Post("/", testController.CreateTest)
		myTest.Post("/:idTest", testController.UpdateTest)
		myTest.Delete("/questions/delete", testQuestionController.DeleteTestQuestion)
		myTest.Post("/:id/questions", testQuestionController.CreateTestQuestion)

	}
	candidates := app.Group("/candidates")
	{
		candidates.Post("/:id", testCandidateController.CreateTestCandidate)
	}
	tests := app.Group("/tests")
	{
		tests.Get("/", testController.FindTests)
	}
	skill := app.Group("/skill")
	{
		skill.Post("/", skillController.CreateSkill)
	}
	skills := app.Group("/skills")
	{
		skills.Get("/", skillController.FindSkills)
	}
	question := app.Group("/questions")
	{
		question.Post("/edit", questionController.CreateQuestion)
		question.Get("/", questionController.FindQuestion)
	}
	candidate := app.Group("/candidate")
	{
		candidate.Post("/", candidateController.CreateCandidate)
	}
	app.Get("/testscandidates", testCandidateController.FindTestsCandidates)
	app.Get("/quiz/:idTestCandidate", testCandidateController.FindQuiz)
	app.Get("/startTest/:idTestCandidate", controllers.StartTest)
	app.Patch("/quiz/status/:idTestCandidate", testCandidateController.UpdateTestStatus)
	app.Patch("/quiz/currentQuestion/:idTestCandidate", testCandidateController.UpdateCurrentQuestion)
	app.Get("/my-tests/:id", testController.GetTest)
	app.Post("/my-tests/clone/:id", testController.CloneTest)

}
