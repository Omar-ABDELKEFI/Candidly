package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTestCandidate(testCandidate models.TestCandidate) (models.TestCandidate, error) {
	log.Println("Creating testCandidate ...")
	db := database.DB

	err := db.Create(&testCandidate).Error

	if err != nil {
		return testCandidate, err
	}
	log.Println("created testQuestion : ", testCandidate)

	return testCandidate, nil
}

//type sum struct {
//	Score      float64 `json:"score"`
//	TestStatus string  `json:"test_status"`
//}

func CalculateScore(idTestCandidate uint64) (models.TestCandidate, error) {
	db := database.DB
	var testCandidate models.TestCandidate
	if err := db.Exec("UPDATE test_candidates AS dest ,"+
		"(SELECT score,CASE WHEN `src`.`score`<src.passing_score THEN 'failed' ELSE 'passed' END AS test_status FROM "+
		"(SELECT ROUND(sum(answers.point * questions.max_points) / sum(max_points)) AS score, tests.passing_score AS passing_score "+
		" FROM `answers` inner join questions on questions.id = answers.question_id AND answers.test_candidate_id = ?"+
		" inner join test_candidates on test_candidates.id =answers.test_candidate_id AND test_candidates.id = ?"+
		" inner join tests on tests.id = test_candidates.test_id GROUP BY `test_candidates`.`id`) AS src ) AS srcScore"+
		" SET `dest`.`score` = `srcScore`.`score`, `dest`.`test_status` = `srcScore`.`test_status`"+
		" WHERE dest.id = ?", idTestCandidate, idTestCandidate, idTestCandidate).Find(&testCandidate).Error; err != nil {
		log.Println("scoreTest", err)

		return testCandidate, err
	}
	log.Println(testCandidate, "tests tests")

	return testCandidate, nil
}
func FindTestsCandidates() ([]models.TestsCandidatesResponse, error) {
	db := database.DB
	var testsCandidates []models.TestsCandidatesResponse
	err := db.Table("test_candidates").
		Joins("INNER JOIN tests on test_candidates.test_id = tests.id").
		Joins("INNER JOIN candidates on test_candidates.candidate_id = candidates.id").
		Select("tests.name as test_name ,candidates.name as candidate_name , candidates.email as candidate_email , test_candidates.score ,test_candidates.test_status, test_candidates.id").
		Find(&testsCandidates).Error
	if err != nil {
		return testsCandidates, err
	}
	log.Println("created testQuestion : ", testsCandidates)

	return testsCandidates, nil

}
func FindQuiz(testId uint64) (models.Test, error) {
	db := database.DB
	var quiz models.Test
	err := db.Table("tests").
		Preload("Questions").
		Preload("Questions.Choices").
		Where("tests.id = ?", testId).
		Find(&quiz).Error

	if err != nil {
		return quiz, err
	}
	log.Println("created testQuestion : ", testId)

	return quiz, nil

}
