package repositories

import (
	"github.com/tekab-dev/tekab-test/database"
	"github.com/tekab-dev/tekab-test/models"
	"log"
)

func CreateTestCandidate(testCandidate models.TestCandidate) (models.TestCandidate, error) {
	log.Println("Creating testCandidate ...")
	db, err := database.GetDb()
	if err != nil {
		return testCandidate, err
	}
	err = db.Create(&testCandidate).Error

	if err != nil {
		return testCandidate, err
	}
	log.Println("created testQuestion : ", testCandidate)

	return testCandidate, nil
}

type sum struct {
	Score      float64 `json:"score"`
	TestStatus string  `json:"test_status"`
}

func CalculateScore(idTestCandidate uint64) (models.TestCandidate, error) {
	db, err := database.GetDb()
	var testCandidate models.TestCandidate
	if err != nil {
		log.Println("error db")
		return testCandidate, err
	}
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
