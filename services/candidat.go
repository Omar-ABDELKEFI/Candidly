package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateCandidat(candidat models.Candidat) (models.Candidat, error) {
	//Create new candidat
	candidat, err := repositories.CreateCandidat(candidat)
	if err != nil {
		return candidat, err
	}
	return candidat, nil
}
