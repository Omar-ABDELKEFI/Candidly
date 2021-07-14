package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTest(test models.Test) (models.Test, error) {
	log.Println("Creating test ...")
	db, err := database.GetDb()
	if err != nil {
		return test, err
	}
	err = db.Create(&test).Error

	if err != nil {
		return test, err
	}
	log.Println("created test : ", test)

	return test, nil
}

func FindTests(skillsID []int64) ([]models.Test, error) {
	db, err := database.GetDb()
	var tests []models.Test
	if err != nil {
		log.Println("error db")
		return nil, err
	}

	if len(skillsID) == 0 {
		if err := db.Table("tests").Find(&tests).Error; err != nil {
			log.Println("question", err)

			return nil, err

		}
		return tests, nil
	}
	log.Println(len(skillsID), "len(skillsID)")
	if err := db.
		Joins("inner JOIN test_questions ON tests.id = test_questions.test_id ").
		Joins("inner JOIN questions ON questions.id = test_questions.question_id AND questions.id IN ?", skillsID).
		Group("tests.id").
		Find(&tests).Error; err != nil {
		log.Println("question", err)

		return nil, err
	}
	log.Println(tests, "tests tests")

	return tests, nil
}
