package models

import "gorm.io/gorm"

type AnswerChoices struct {
	gorm.Model
	AnswerId  uint64 `json:"answer_id"`
	ChoicesId uint64 `json:"choices_id" validate:"required"`
}

type QuestionChoices struct {
	QuestionId uint64
	Type       string
	Points     int
	IsAnswer   bool
	ChoiceId   uint64
}
