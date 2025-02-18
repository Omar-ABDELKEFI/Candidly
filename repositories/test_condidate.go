package repositories

import (
	"github.com/dev/test/database"
	"github.com/dev/test/models"
	"log"
)

func CreateTestCandidate(testCandidates []models.TestCandidate) ([]string, []models.TestCandidate, []string) {
	log.Println("Creating testCandidate ...")
	thereError := false
	var emails []string
	var candidateId []uint64
	var errors []string
	db := database.DB
	for _, testCandidate := range testCandidates {
		if err := db.Create(&testCandidate).Error; err != nil {
			thereError = true
			candidateId = append(candidateId, testCandidate.CandidateID)
			errors = append(errors, err.Error())
		}
	}
	log.Println(errors, "errors in repositories/test_candidate (CreateTestCandidate) ")
	if thereError == true {
		db.Table("candidates").Select("email").Where("id IN ?", candidateId).Find(&emails)
		return emails, testCandidates, errors
	}

	log.Println("created testQuestion : ", testCandidates)

	return emails, testCandidates, errors
}

//type sum struct {
//	Score      float64 `json:"score"`
//	TestStatus string  `json:"test_status"`
//}

func CalculateScore(candidateId uint64, testId uint64) (models.TestCandidate, error) {
	db := database.DB
	var testCandidate models.TestCandidate
	if err := db.Exec("UPDATE test_candidates AS dest ,"+
		"(SELECT score,CASE WHEN `src`.`score`<src.passing_score THEN 'failed' ELSE 'passed' END AS test_status FROM "+
		"(SELECT ROUND(sum(answers.point * questions.max_points) / sum(max_points)) AS score, tests.passing_score AS passing_score "+
		" FROM `answers` inner join questions on questions.id = answers.question_id AND answers.candidate_id = ? AND answers.test_id = ?"+
		" inner join test_candidates on test_candidates.candidate_id =answers.candidate_id AND test_candidates.test_id = answers.test_id AND test_candidates.candidate_id = ? AND test_candidates.test_id = ?"+
		" inner join tests on tests.id = test_candidates.test_id GROUP BY `test_candidates`.`candidate_id` , `test_candidates`.`test_id`) AS src ) AS srcScore"+
		" SET `dest`.`score` = `srcScore`.`score`, `dest`.`test_status` = `srcScore`.`test_status`"+
		" WHERE dest.candidate_id = ? AND dest.test_id = ?", candidateId, testId, candidateId, testId, candidateId, testId).Where("candidate_id = ? and test_id = ?", candidateId, testId).Find(&testCandidate).Error; err != nil {
		log.Println("scoreTest", err)

		return testCandidate, err
	}
	log.Println(testCandidate, "tests tests")

	return testCandidate, nil
}
func FindTestsCandidates(testId string, isNil bool) ([]models.TestsCandidatesResponse, error) {
	var conditionJoin string
	if isNil {
		conditionJoin = ""
	} else {
		conditionJoin = " and test_candidates.test_id = " + testId

	}
	log.Println("conditionJoin : ", conditionJoin)
	db := database.DB

	var testsCandidates []models.TestsCandidatesResponse
	err := db.Table("test_candidates").
		Joins("INNER JOIN tests on test_candidates.test_id = tests.id" + conditionJoin).
		Joins("INNER JOIN candidates on test_candidates.candidate_id = candidates.id").
		Select("tests.name as test_name ,candidates.name as candidate_name , candidates.email as candidate_email , test_candidates.score ,test_candidates.test_status , CONCAT(test_candidates.test_id, '-' ,test_candidates.candidate_id) as test_candidate_id").
		Find(&testsCandidates).Error
	if err != nil {
		return testsCandidates, err
	}
	log.Println("created testQuestion : ", testsCandidates)

	return testsCandidates, nil

}
func StartTest(testId uint64, candidateId uint64) (models.StartTest, error) {
	var results models.StartTest
	db := database.DB
	raw := db.Table("tests").Select("tests.name as name", "test_candidates.created_at as created_at", "test_candidates.test_status as test_status", "test_candidates.score as score ", "candidates.email as email", "tests.time_limit as time_limit", "test_candidates.current_question as current_question, test_candidates.updated_at as updated_at").Where("tests.id = ?", testId).Joins("inner join candidates on candidates.id= ?", candidateId).
		Joins("inner join test_candidates on test_candidates.test_id = ? AND test_candidates.candidate_id = ?", testId, candidateId).Row()
	if raw.Err() != nil {

	}

	raw.Scan(&results.Name, &results.CreatedAt, &results.TestStatus, &results.Score, &results.Email, &results.TimeLimit, &results.CurrentQuestion, &results.UpdatedAt)
	rows, _ := db.Table("test_questions").Select("questions.name", "questions.type", "questions.expected_time").Where("test_id = ?", testId).Joins("inner join questions on test_questions.question_id=questions.id ").Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &results.Questions)

		// do something
	}

	return results, nil
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

func UpdateTestStatus(testId uint64, candidateId uint64, testStatus models.UpdateTestStatus) (models.UpdateTestStatusOutput, error) {
	db := database.DB
	var quiz models.TestCandidate
	var testStatusOutput models.UpdateTestStatusOutput
	err := db.Find(&quiz, "test_id = ? AND candidate_id = ?", testId, candidateId).Error
	if err != nil {
		return testStatusOutput, err
	}
	if quiz.TestStatus != "canceled" && quiz.TestStatus == "waiting" {
		db.Model(&quiz).Update("test_status", testStatus.TestStatus)
	}

	testStatusOutput.TestStatus = quiz.TestStatus
	testStatusOutput.UpdatedAt = quiz.UpdatedAt

	return testStatusOutput, nil

}
func UpdateCurrentQuestion(testId uint64, candidateId uint64, currentQuestion models.UpdateCurrentQuestion) (models.UpdateCurrentQuestionOutput, error) {
	db := database.DB
	var quiz models.TestCandidate
	var currentQuestionOutput models.UpdateCurrentQuestionOutput

	err := db.Find(&quiz, "test_id = ? AND candidate_id = ?", testId, candidateId).Error
	if err != nil {
		return currentQuestionOutput, err
	}
	db.Model(&quiz).Update("current_question", currentQuestion.CurrentQuestion)
	currentQuestionOutput.CurrentQuestion = quiz.CurrentQuestion
	currentQuestionOutput.UpdatedAt = quiz.UpdatedAt
	return currentQuestionOutput, nil

}
