package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateCandidate(candidat models.Candidate) (models.Candidate, error) {
	log.Println("Creating Candidate ...")
	db, err := database.GetDb()
	if err != nil {
		return candidat, err
	}
	err = db.Create(&candidat).Error

	if err != nil {
		return candidat, err
	}
	log.Println("created candidate : ", candidat)

	return candidat, nil
}
