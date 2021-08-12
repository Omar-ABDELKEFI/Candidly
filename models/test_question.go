package models

type TestQuestion struct {
	OwnModel
	QuestionID uint64 `json:"question_id" validate:"required"`
	TestID     uint64 `json:"test_id"`
}
