package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateCandidate(candidate models.Candidate) (models.Candidate, error) {
	//Create new candidate
	candidate, err := repositories.CreateCandidate(candidate)
	if err != nil {
		return candidate, err
	}
	return candidate, nil
}
