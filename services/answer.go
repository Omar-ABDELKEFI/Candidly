package services

import (
	"errors"
	"github.com/dev/test/common"
	"github.com/dev/test/models"
	"github.com/dev/test/repositories"
	"log"
)

//Create new answer
func CreateAnswer(answer models.Answer) (models.Answer, error) {
	// Check whether the answer is related to mcq or mca question
	if answer.AnswerChoices != nil {
		// get all choices related to the answered question
		questionChoices, err := repositories.GetQuestionChoices(answer)
		if err != nil {
			return answer, err
		}
		//calculate score for mca questions

		//count the correct answers then calculate the score for correct answer and the coast of wrong answer
		//for every choice choosen by the candidate if it is correct then add the correct score value to the question score
		//else soustract the wrong answer coast from the question score
		if questionChoices[0].Type == "mca" {
			//calculate correct answer score
			correctAnswerCount := 0
			for i := 0; i < len(questionChoices); i++ {
				if questionChoices[i].IsAnswer {
					correctAnswerCount++
				}
			}
			correctAnswerScore := float64(100) / float64(correctAnswerCount)
			log.Println("correctAnswerScore", correctAnswerScore)
			//calculate Wrong answer coast
			wrongAnswerScore := float64(100) / float64(len(questionChoices)-correctAnswerCount)
			//calculate score
			for i := 0; i < len(answer.AnswerChoices); i++ {
				// getting current choice index
				choiceIndex := common.ChoiceIndex(questionChoices, answer.AnswerChoices[i].ChoicesID)
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
			choiceIndex := common.ChoiceIndex(questionChoices, answer.AnswerChoices[0].ChoicesID)
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
