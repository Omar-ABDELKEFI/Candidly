package services

import (
	"github.com/dev/test/models"
	"github.com/dev/test/repositories"
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

	tests, err := repositories.FindTests(tabSkillsInteger)
	if err != nil {
		log.Println("error db")
		return nil, err
	}

	return tests, nil
}

func GetMyTests() ([]models.MyTests, error) {
	//Create new test
	newTest, err := repositories.GetMyTests()
	if err != nil {
		return nil, err
	}
	return newTest, err
}

func GetTest(idTest uint64) (models.Test, error) {
	test, err := repositories.GetTest(idTest)
	if err != nil {
		return test, err
	}
	return test, err
}
func CloneTest(idTest uint64, input models.CloneTestInput) (models.MyTests, error) {
	clonedTest, err := repositories.CloneTest(idTest, input)
	if err != nil {
		return clonedTest, err
	}
	return clonedTest, err
}
