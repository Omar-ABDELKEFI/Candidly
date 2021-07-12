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
