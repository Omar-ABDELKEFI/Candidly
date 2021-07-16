package models

import "gorm.io/gorm"

type AnswerChoices struct {
	gorm.Model
	AnswerId  uint64 `json:"answer_id" validate:"required"`
	ChoicesId uint64 `json:"choices_id" validate:"required"`
}
