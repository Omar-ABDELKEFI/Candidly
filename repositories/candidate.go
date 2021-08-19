package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateCandidate(candidate []models.Candidate) ([]models.Candidate, error) {
	log.Println("Creating Candidate ...")
	db := database.DB
	db.Model(&candidate).Association("test_candidate").Append()
	err := db.Create(&candidate).Error

	if err != nil {
		return candidate, err
	}
	log.Println("created candidate : ", candidate)

	return candidate, nil
}
