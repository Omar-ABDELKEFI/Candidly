package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateCandidate(input models.CreateCandidateInput) (models.Candidate, error) {
	//Create new candidate
	candidate := models.NewCandidate(input)
	newCandidate, err := repositories.CreateCandidate(candidate)
	if err != nil {
		return candidate, err
	}
	return newCandidate, nil
}
