package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTestQuestion(testQuestion models.TestQuestion) (models.TestQuestion, error) {
	log.Println("Creating testQuestion ...")
	db := database.DB

	err := db.Create(&testQuestion).Error

	if err != nil {
		return testQuestion, err
	}
	log.Println("created testQuestion : ", testQuestion)

	return testQuestion, nil
}
func DeleteTestQuestion(id uint64) error {
	log.Println("Deleting testQuestion ...")
	db := database.DB
	var testQuestion models.TestQuestion
	err := db.Table("test_questions").Where("id = ?", id).First(&testQuestion).Error
	if err != nil {
		return err
	}
	err = db.Delete(&testQuestion).Error
	if err != nil {
		return err
	}

	return nil
}
