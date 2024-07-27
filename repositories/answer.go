package repositories

import (
	"github.com/dev/test/database"
	"github.com/dev/test/models"
	"log"
)

//get all choices of a question
func GetQuestionChoices(answer models.Answer) ([]models.QuestionChoices, error) {
	log.Println("getting question choices")
	var questionChoices []models.QuestionChoices
	db := database.DB

	err := db.Table("questions").
		Select("questions.id as question_id, questions.type, choices.is_answer, choices.id as choice_id").
		Joins("inner join choices on choices.question_id = questions.id").
		Where("choices.question_id = (?)", answer.QuestionID).
		Find(&questionChoices).Error
	if err != nil {
		return questionChoices, err
	}
	return questionChoices, err

}

//create an answer
func CreateAnswer(answer models.Answer) (models.Answer, error) {

	log.Println("Creating Answer ...")
	db := database.DB

	err := db.Create(&answer).Error
	if err != nil {
		return answer, err
	}

	return answer, nil
}
