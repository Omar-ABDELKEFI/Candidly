package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
)

func CreateQuestion(input models.CreateQuestionInput) (models.Question, error) {
	question := models.NewQuestion(input)
	newQuestion, err := repositories.CreateQuestion(question)
	log.Println("test")
	if err != nil {
		return question, err
	}
	return newQuestion, err
}
func FindQuestions(sort []string, difficulty []string) ([]models.Question, error) {
	questions, err := repositories.FindQuestions(sort, difficulty)

	if err != nil {
		log.Println("error db")
		return nil, err
	}

	return questions, nil
}
func GetQuestion(idQuestion uint64) (models.Question, error) {
	question, err := repositories.GetQuestion(idQuestion)
	if err != nil {
		return question, err
	}
	return question, err
}
func UpdateQuestion(idQuestion uint64, input models.CreateQuestionInput) (models.Question, error) {
	log.Println("inputtt", input)
	question := models.NewQuestion(input)
	newQuestion, err := repositories.UpdateQuestion(idQuestion, question)
	log.Println("test")
	if err != nil {
		return question, err
	}
	return newQuestion, err
}
