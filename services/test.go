package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

func CreateTest(test models.Test) (models.Test, error) {
	//Create new test
	newTest, err := repositories.CreateTest(test)
	if err != nil {
		return test, err
	}
	return newTest, err
}
