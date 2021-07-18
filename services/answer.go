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
	if answer.AnswerChoices != nil {
		questionChoices, err := repositories.GetQuestionChoices(answer)
		if err != nil {
			return answer, err
		}
		if questionChoices[0].Type == "mca" {
			//calculate correct answer score
			correctAnswerCount := 0
			for i := 0; i < len(questionChoices); i++ {
				if questionChoices[i].IsAnswer {
					correctAnswerCount++
				}
			}
			correctAnswerScore := float64(100) / float64(correctAnswerCount)
			//calculate correct answer score
			wrongAnswerScore := float64(100) / float64(len(questionChoices)-correctAnswerCount)
			//calculate score
			for i := 0; i < len(answer.AnswerChoices); i++ {
				// getting current choice index
				choiceIndex := common.ChoiceIndex(questionChoices, answer.AnswerChoices[i].ChoicesId)
				if choiceIndex == -1 {
					log.Println("choice not found")
					indexError := errors.New("could not find choice")
					return answer, indexError
				}
				if questionChoices[choiceIndex].IsAnswer {
					answer.Point += correctAnswerScore
				} else {
					answer.Point -= wrongAnswerScore
				}
			}
			if answer.Point < 0 {
				answer.Point = 0
			}
		} else {
			// calculate mcq score
			choiceIndex := common.ChoiceIndex(questionChoices, answer.AnswerChoices[0].ChoicesId)
			if choiceIndex == -1 {
				log.Println("choice not found")
				indexError := errors.New("could not find choice")
				return answer, indexError
			}
			if questionChoices[choiceIndex].IsAnswer {
				answer.Point = 100
			}
		}
	}

	newAnswer, err := repositories.CreateAnswer(answer)
	if err != nil {
		return answer, err
	}
	return newAnswer, err
}
