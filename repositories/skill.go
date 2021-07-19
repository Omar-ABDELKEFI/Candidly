package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"gorm.io/gorm/clause"
)

func CreateSkill(skill models.Skill) error {
	db := database.DB

	var i uint
	for i = 0; i < 10000; i++ {
		skill.ID = i + 25000
		db.Omit(clause.Associations).Create(&skill)
	}
	//if err := db.Omit(clause.Associations).Create(&skill).Error; err != nil {
	//	log.Println("skill", err)
	//
	//	return err
	//
	//}

	return nil
}
