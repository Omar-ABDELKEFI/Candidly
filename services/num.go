package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
)

func CreateNum(question models.Question) (err error) {
	err = repositories.CreateNum(question)
	log.Println("test")
	if err != nil {
		return err
	}
	return nil
}
