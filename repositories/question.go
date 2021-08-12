package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/gorm"
	"log"
)

func CreateQuestion(question models.Question) (models.Question, error) {
	db := database.DB

	if err := db.Create(&question).Error; err != nil {
		log.Println("question", err)
		return question, err
	}
	return question, nil
}

func FindQuestion(sort []string, difficulty []string) ([]models.Question, error) {
	db := database.DB
	var question []models.Question

	/*log.Println(sort, len(sort))
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
	*/
	err := db.Table("questions").Preload("TestQuestions", func(db *gorm.DB) *gorm.DB {
		return db.Select("test_questions.id , test_questions.question_id , test_questions.test_id")
	}).Preload("Skill", func(db *gorm.DB) *gorm.DB {
		return db.Select("skills.id , skills.name")
	}).Find(&question).Error
	if err != nil {
		log.Println("question", err)

		return nil, err
	}

	return question, nil
}
