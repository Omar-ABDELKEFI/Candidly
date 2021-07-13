package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateCandidat(candidat models.Candidat) (models.Candidat, error) {
	log.Println("Creating Candidat ...")
	db, err := database.GetDb()
	if err != nil {
		return candidat, err
	}
	err = db.Create(&candidat).Error

	if err != nil {
		return candidat, err
	}
	log.Println("created candidat : ", candidat)

	return candidat, nil
}
