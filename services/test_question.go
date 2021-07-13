package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateTestQuestion(testQuestion models.TestQuestion) (models.TestQuestion, error) {
	//Create new testQuestion
	newTest, err := repositories.CreateTestQuestion(testQuestion)
	if err != nil {
		return testQuestion, err
	}
	return newTest, err
}
