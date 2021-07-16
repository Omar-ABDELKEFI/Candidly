package models

import "gorm.io/gorm"

type TestQuestion struct {
	gorm.Model
	QuestionId uint64 `json:"question_id" validate:"required"`
	TestId     uint64 `json:"test_id" validate:"required"`
}
