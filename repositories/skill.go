package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/gorm/clause"
	"log"
)

func CreateSkill(skill models.Skill) error {
	db, err := database.GetDb()

	if err != nil {
		log.Println("err db")
		return err
	}

	if err := db.Omit(clause.Associations).Create(&skill).Error; err != nil {
		log.Println("skill", err)

		return err

	}

	return nil
}
