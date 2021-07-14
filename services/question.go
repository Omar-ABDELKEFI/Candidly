package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
)

func CreateQuestion(question models.Question) (err error) {
	err = repositories.CreateQuestion(question)
	log.Println("test")
	if err != nil {
		return err
	}
	return nil
}
func FindQuestion(sort []string, difficulty []string) ([]models.Question, error) {
	questions, err := repositories.FindQuestion(sort, difficulty)

	if err != nil {
		log.Println("error db")
		return nil, err
	}

	return questions, nil
}
