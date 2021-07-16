package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
)

//Create new answer
func CreateAnswer(answer models.Answer) (models.Answer, error) {
	newAnswer, err := repositories.CreateAnswer(answer)
	if err != nil {
		return answer, err
	}
	return newAnswer, err

}
