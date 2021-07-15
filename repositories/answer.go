package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func GetQuestionChoices(answer models.Answer) ([]models.AnswerDbResponse, error) {
	log.Println("Getting Question choices ...")
	db, err := database.GetDb()
	if err != nil {
		return nil, err
	}
	var response []models.AnswerDbResponse
	err = db.Table("questions").
		Select("questions.id, questions.type, questions.points, choices.is_answer, choices.id as choice_id").
		Joins("inner join choices on choices.question_id = questions.id").
		Where("choices.question_id = (?)", db.Table("choices").Select("question_id").Where("id = ?", answer.ChoicesId)).Find(&response).Error
	log.Println("response ", response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func CreateAnswer(answer models.Answer) (models.Answer, error) {

	log.Println("Creating Answer ...")
	db, err := database.GetDb()
	if err != nil {
		return answer, err
	}

	err = db.Create(&answer).Error
	if err != nil {
		return answer, err
	}

	return answer, nil
}
