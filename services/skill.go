package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
)

func CreateSkill(input models.CreateSkillInput) (models.Skill, error) {
	//Create new test
	skill := models.NewSkill(input)
	newSkill, err := repositories.CreateSkill(skill)
	log.Println("skill")
	if err != nil {
		return skill, err
	}
	return newSkill, nil
}
func FindSkills() ([]models.Skill, error) {
	var skills []models.Skill
	//Get all skills
	skills, err := repositories.FindSkills()
	if err != nil {
		return skills, err
	}
	return skills, nil
}
