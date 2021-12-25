package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateTestCandidate(testCandidate []models.TestCandidate) ([]string, []models.TestCandidate, []string) {
	//Create new testCandidate
	emailsDuplicate, newTestCandidate, errDuplicate := repositories.CreateTestCandidate(testCandidate)
	if errDuplicate != nil {
		return emailsDuplicate, newTestCandidate, errDuplicate
	}
	return nil, newTestCandidate, nil
}

func CalculateScore(candidateId uint64, testId uint64) (models.TestCandidate, error) {
	var TestCandidate models.TestCandidate
	TestCandidate, err := repositories.CalculateScore(candidateId, testId)
	if err != nil {
		return TestCandidate, err
	}
	return TestCandidate, err
}
func FindTestsCandidates(testId string, isNil bool) ([]models.TestsCandidatesResponse, error) {
	var testsCandidates []models.TestsCandidatesResponse
	testsCandidates, err := repositories.FindTestsCandidates(testId, isNil)
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
func UpdateTestStatus(testId uint64, candidateId uint64, testStatus models.UpdateTestStatus) (models.UpdateTestStatusOutput, error) {
	testStatusOutput, err := repositories.UpdateTestStatus(testId, candidateId, testStatus)
	if err != nil {
		return testStatusOutput, err
	}
	return testStatusOutput, nil
}

func UpdateCurrentQuestion(testId uint64, candidateId uint64, currentQuestion models.UpdateCurrentQuestion) (models.UpdateCurrentQuestionOutput, error) {
	currentQuestionOutput, err := repositories.UpdateCurrentQuestion(testId, candidateId, currentQuestion)
	if err != nil {
		return currentQuestionOutput, err
	}
	return currentQuestionOutput, nil
}
