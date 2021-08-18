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

func CalculateScore(candidateId uint64, testId uint64) (models.TestCandidate, error) {
	var TestCandidate models.TestCandidate
	TestCandidate, err := repositories.CalculateScore(candidateId, testId)
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
func StartTest(testId uint64, candidateId uint64) (models.StartTest, error) {
	results, err := repositories.StartTest(testId, candidateId)
	if err != nil {
		return results, err
	}
	return results, nil
}

func FindQuiz(testID uint64) (models.Test, error) {
	var quiz models.Test
	quiz, err := repositories.FindQuiz(testID)
	if err != nil {
		return quiz, err
	}
	return quiz, nil
}
func UpdateTestStatus(testId uint64, candidateId uint64, testStatus models.UpdateTestStatus) (models.UpdateTestStatus, error) {
	testStatus, err := repositories.UpdateTestStatus(testId, candidateId, testStatus)
	if err != nil {
		return testStatus, err
	}
	return testStatus, nil
}

func UpdateCurrentQuestion(testId uint64, candidateId uint64, currentQuestion models.UpdateCurrentQuestion) (models.UpdateCurrentQuestion, error) {
	currentQuestion, err := repositories.UpdateCurrentQuestion(testId, candidateId, currentQuestion)
	if err != nil {
		return currentQuestion, err
	}
	return currentQuestion, nil
}
