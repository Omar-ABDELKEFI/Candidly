package models

type TestQuestion struct {
	OwnModel
	QuestionId uint64 `json:"question_id" validate:"required"`
	TestId     uint64 `json:"test_id" validate:"required"`
}
