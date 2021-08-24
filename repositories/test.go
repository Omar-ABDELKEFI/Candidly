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

func GetMyTests() ([]models.MyTests, error) {

	var getTest []models.MyTests
	db := database.DB
	if err := db.Raw("SELECT SUM(questions.expected_time) as expected_time ," +
		" COUNT(DISTINCT test_questions.question_id) number_question," +
		" tests.id AS test_id," +
		" tests.name AS test_name," +
		" (SELECT COUNT(test_candidates.candidate_id) FROM test_candidates WHERE test_candidates.test_id=tests.id) AS number_candidate" +
		" FROM" +
		" test_questions" +
		" RIGHT JOIN tests ON tests.id = test_questions.test_id" +
		" LEFT JOIN questions ON questions.id = test_questions.question_id" +
		" GROUP BY" +
		" tests.id;").
		Scan(&getTest).
		Error; err != nil {
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

func GetTest(testId uint64) (models.Test, error) {
	db := database.DB
	var test models.Test

	if err := db.
		Table("tests").Preload("Questions").Preload("Questions.Skill").Where("tests.id = ?", testId).
		Find(&test).Error; err != nil {
		log.Println("GetTest ", err)

		return test, err
	}

	return test, nil
}
