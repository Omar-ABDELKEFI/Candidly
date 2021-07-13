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

//func FindQuestion(sort string) ([]models.Question,error) {
//	db, err := database.GetDb()
//
//	if err != nil {
//		log.Println("error db")
//		return nil,err
//	}
//
//	if err := db.Table("questions").Where("ty"); err != nil {
//		log.Println("question", err)
//
//		return err
//
//	}
//
//	return nil
//}
