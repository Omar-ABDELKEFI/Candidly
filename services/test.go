package services

import (
	"github.com/tekab-dev/tekab-test/models"
	"github.com/tekab-dev/tekab-test/repositories"
	"log"
	"strconv"
	"strings"
)

func CreateTest(test models.Test) (models.Test, error) {
	//Create new test

	newTest, err := repositories.CreateTest(test)
	if err != nil {
		return test, err
	}
	return newTest, err
}
func UpdateTest(test models.Test, idTest uint64) (models.Test, error) {
	//Create new test
	newTest, err := repositories.UpdateTest(test, idTest)
	if err != nil {
		return test, err
	}
	return newTest, err
}
func FindTests(skillsId string) ([]models.Test, error) {
	var tabSkillsInteger []int64
	var tabSkillsString []string
	if len(skillsId) != 0 {
		tabSkillsString = strings.Split(skillsId, ",")
	}
	log.Println(len(tabSkillsString), "len(tabSkillsString)")
	for _, i := range tabSkillsString {
		j, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			panic(err)
		}
		tabSkillsInteger = append(tabSkillsInteger, j)
	}
	log.Println(tabSkillsInteger, "ssssssssssss")
	tests, err := repositories.FindTests(tabSkillsInteger)
	if err != nil {
		log.Println("error db")
		return nil, err
	}

	return tests, nil
}
