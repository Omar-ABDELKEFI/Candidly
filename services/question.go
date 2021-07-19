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
func FindQuestion(sort []string, difficulty []string) ([]models.Question, error) {
	questions, err := repositories.FindQuestion(sort, difficulty)

	if err != nil {
		log.Println("error db")
		return nil, err
	}

	return questions, nil
}
