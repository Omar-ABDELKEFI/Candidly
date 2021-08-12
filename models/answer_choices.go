package models

import "gorm.io/gorm"

type AnswerChoices struct {
	gorm.Model
	AnswerID  uint64 `json:"answer_id"`
	ChoicesID uint64 `json:"choices_id" validate:"required"`
}

type QuestionChoices struct {
	QuestionID uint64
	Type       string
	IsAnswer   bool
	ChoiceID   uint64
}
