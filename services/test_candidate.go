package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateTestCandidate(testCandidate models.TestCandidate) (models.TestCandidate, error) {
	//Create new testCandidate
	newTestCandidate, err := repositories.CreateTestCandidate(testCandidate)
	if err != nil {
		return newTestCandidate, err
	}
	return newTestCandidate, err
}
