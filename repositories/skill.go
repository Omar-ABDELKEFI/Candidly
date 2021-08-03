package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateSkill(skill models.Skill) (models.Skill, error) {
	db := database.DB
	err := db.Create(&skill).Error
	if err != nil {
		log.Println("skill", err)
		return skill, err
	}
	return skill, nil
}
func FindSkills() (skills []models.Skill, err error) {
	db := database.DB
	//Find skills
	if err := db.Table("skills").Find(&skills).Error; err != nil {
		return skills, err
	}
	return skills, nil
}
func FindOrCreateSkill(skillName string) (uint, error) {
	db := database.DB
	var skill models.Skill
	err := db.FirstOrCreate(&skill, models.Skill{Name: skillName}).Error
	if err != nil {
		return 0, err
	}
	return skill.ID, err
}
