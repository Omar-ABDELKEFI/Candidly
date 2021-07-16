package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

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
