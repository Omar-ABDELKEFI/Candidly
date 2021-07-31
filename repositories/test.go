package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTest(test models.Test) (models.Test, error) {
	log.Println("Creating test ...")
	db := database.DB

	err := db.Create(&test).Error

	if err != nil {
		return test, err
	}
	log.Println("created test : ", test)

	return test, nil
}
func UpdateTest(test models.Test, idTest uint64) (models.Test, error) {
	log.Println("Creating test ...")
	var getTest models.Test
	db := database.DB
	if err := db.Model(&getTest).Where("id = ?", idTest).Updates(&test).First(&getTest, idTest).Error; err != nil {
		return getTest, err
	}

	return getTest, nil
}

func FindTests(skillsID []int64) ([]models.Test, error) {
	db := database.DB
	var tests []models.Test
	if len(skillsID) == 0 {
		if err := db.Table("tests").Find(&tests).Error; err != nil {
			log.Println("tests", err)

			return nil, err

		}
		return tests, nil
	}
	if err := db.
		Joins("inner JOIN test_questions ON tests.id = test_questions.test_id ").
		Joins("inner JOIN questions ON questions.id = test_questions.question_id AND questions.id IN ?", skillsID).
		Group("tests.id").
		Find(&tests).Error; err != nil {
		log.Println("tests", err)

		return nil, err
	}
	log.Println(tests, "tests tests")

	return tests, nil
}
