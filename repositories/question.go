package repositories

import (
	"github.com/dev/test/database"
	"github.com/dev/test/models"
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

func FindQuestions(sort []string, difficulty []string) ([]models.Question, error) {
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
	err := db.Table("questions").Preload("Choices").Preload("TestQuestions", func(db *gorm.DB) *gorm.DB {
		return db.Select("test_questions.id , test_questions.question_id , test_questions.test_id")
	}).Preload("Skill", func(db *gorm.DB) *gorm.DB {
		return db.Select("skills.id , skills.name")
	}).Order("questions.id DESC").
		Find(&question).Error
	if err != nil {
		log.Println("question", err)

		return nil, err
	}

	return question, nil
}
func GetQuestion(questionId uint64) (models.Question, error) {
	db := database.DB
	var question models.Question

	err := db.Table("questions").Preload("Choices").
		Preload("Skill", func(db *gorm.DB) *gorm.DB {
			return db.Select("skills.id , skills.name")
		}).Order("questions.id DESC").
		Where("questions.id = ?", questionId).
		Find(&question).Error
	if err != nil {
		log.Println("question", err)
		return question, err

	}
	return question, nil
}
func UpdateQuestion(questionId uint64, question models.Question) (models.Question, error) {
	db := database.DB
	log.Println("quessssss", question)
	question.ID = uint(questionId)
	err := db.Delete(models.Choices{}, "question_id = ?", questionId).Error
	if err != nil {
		return question, err
	}
	err = db.Model(&question).Where("id = ?", questionId).Updates(&question).Error
	if err != nil {
		return question, err
	}
	return question, nil
}
