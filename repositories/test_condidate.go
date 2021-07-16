package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTestCandidate(testCandidate models.TestCandidate) (models.TestCandidate, error) {
	log.Println("Creating testCandidate ...")
	db, err := database.GetDb()
	if err != nil {
		return testCandidate, err
	}
	err = db.Create(&testCandidate).Error

	if err != nil {
		return testCandidate, err
	}
	log.Println("created testQuestion : ", testCandidate)

	return testCandidate, nil
}
