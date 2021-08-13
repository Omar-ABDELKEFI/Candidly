package models

import "time"

type TestQuestion struct {
	ID         uint      `gorm:"primarykey"`
	CreatedAt  time.Time `swaggerignore:"true"`
	UpdatedAt  time.Time `swaggerignore:"true"`
	QuestionID uint64    `json:"question_id" validate:"required"`
	TestID     uint64    `json:"test_id"`
}
