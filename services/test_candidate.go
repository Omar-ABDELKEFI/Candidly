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

func CalculateScore(idTestCandidate uint64) (models.TestCandidate, error) {
	var TestCandidate models.TestCandidate
	TestCandidate, err := repositories.CalculateScore(idTestCandidate)
	if err != nil {
		return TestCandidate, err
	}
	return TestCandidate, err
}
func FindTestsCandidates() ([]models.TestsCandidatesResponse, error) {
	var testsCandidates []models.TestsCandidatesResponse
	testsCandidates, err := repositories.FindTestsCandidates()
	if err != nil {
		return testsCandidates, err
	}
	return testsCandidates, nil
}

func FindQuiz(testID uint64) (models.Test, error) {
	var quiz models.Test
	quiz, err := repositories.FindQuiz(testID)
	if err != nil {
		return quiz, err
	}
	return quiz, nil
}
