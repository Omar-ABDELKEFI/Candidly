package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateQuestion(question models.Question) error {
	db, err := database.GetDb()

	if err != nil {
		log.Println("error db")
		return err
	}

	if err := db.Create(&question).Error; err != nil {
		log.Println("question", err)

		return err

	}

	return nil
}

func FindQuestion(sort []string, difficulty []string) ([]models.Question, error) {
	db, err := database.GetDb()
	var question []models.Question
	if err != nil {
		log.Println("error db")
		return nil, err
	}
	log.Println(sort, len(sort))
	if len(sort) == 0 && len(difficulty) == 0 {
		if err := db.Table("questions").Find(&question).Error; err != nil {
			log.Println("question", err)

			return nil, err

		}
		return question, nil
	}
	if err := db.Table("questions").Where("type IN ?", sort).Where("difficulty IN ?", difficulty).Find(&question).Error; err != nil {
		log.Println("question", err)

		return nil, err
	}

	return question, nil
}
