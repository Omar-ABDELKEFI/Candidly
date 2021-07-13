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
