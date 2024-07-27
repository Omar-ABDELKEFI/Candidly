package services

import (
	"github.com/dev/test/models"
	"github.com/dev/test/repositories"
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
func FindOrCreateSkill(skillName string) (uint, error) {
	skillId, err := repositories.FindOrCreateSkill(skillName)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return skillId, nil
}
func FindSkills() ([]models.SkillsResponse, error) {
	var skills []models.SkillsResponse
	//Get all skills
	skills, err := repositories.FindSkills()
	if err != nil {
		return skills, err
	}
	return skills, nil
}
