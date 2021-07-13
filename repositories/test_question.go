package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTestQuestion(testQuestion models.TestQuestion) (models.TestQuestion, error) {
	log.Println("Creating testQuestion ...")
	db, err := database.GetDb()
	if err != nil {
		return testQuestion, err
	}
	err = db.Create(&testQuestion).Error

	if err != nil {
		return testQuestion, err
	}
	log.Println("created testQuestion : ", testQuestion)

	return testQuestion, nil
}
