package services

import (
	"errors"
	"github.com/tekab-dev/tekab-test/common"
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
)

//Create new answer
func CreateAnswer(answer models.Answer) (models.Answer, error) {
	if answer.ChoicesId != 0 {
		// getting question choices
		var questionChoices []models.AnswerDbResponse
		questionChoices, err := repositories.GetQuestionChoices(answer)
		if err != nil {
			return answer, err
		}
		// searching for my current choice index
		choiceIndex := common.ChoiceIndex(questionChoices, answer.ChoicesId)
		if choiceIndex == -1 {
			log.Println("choice not found")
			indexError := errors.New("could not find choice")
			return answer, indexError
		}
		//calculating score for mca questions
		if questionChoices[0].Type == "mca" && questionChoices[choiceIndex].IsAnswer {
			answer.Point = float64(questionChoices[choiceIndex].Points) / float64(len(questionChoices))
		}
		//calculating score for mcq questions
		if questionChoices[0].Type == "mcq" && questionChoices[choiceIndex].IsAnswer {
			answer.Point = float64(questionChoices[choiceIndex].Points)
		}
	}
	newAnswer, err := repositories.CreateAnswer(answer)
	if err != nil {
		return answer, err
	}
	return newAnswer, err
}
