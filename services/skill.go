package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
)

func CreateSkill(skill models.Skill) (err error) {
	err = repositories.CreateSkill(skill)
	log.Println("test")
	if err != nil {
		return err
	}
	return nil
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
