package services

import (
	"github.com/tekab-dev/tekab-test/common"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateCandidate(candidate []models.Candidate) ([]models.Candidate, error) {
	//Create new candidate
	newCandidates, err := repositories.CreateCandidate(candidate)
	common.Send(newCandidates)
	if err != nil {
		return candidate, err
	}
	return newCandidates, nil
}
