package services

import (
	"github.com/dev/test/models"
	"github.com/dev/test/repositories"
)

func CreateTestQuestion(testQuestion models.TestQuestion) (models.TestQuestion, error) {
	//Create new testQuestion
	newTest, err := repositories.CreateTestQuestion(testQuestion)
	if err != nil {
		return testQuestion, err
	}
	return newTest, err
}
func DeleteTestQuestion(testId uint64, questionId uint64) error {
	//Delete testQuestion
	err := repositories.DeleteTestQuestion(testId, questionId)
	if err != nil {
		return err
	}
	return nil
}
