package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/gorm/clause"
	"log"
)

func CreateSkill(skill models.Skill) error {
	db := database.DB

	if err := db.Omit(clause.Associations).Create(&skill).Error; err != nil {
		log.Println("skill", err)

		return err

	}

	return nil
}
func FindSkills() (skills []models.Skill, err error) {
	db := database.DB
	//Find books
	if err := db.Table("skills").Find(&skills).Error; err != nil {
		return skills, err
	}
	return skills, nil
}
