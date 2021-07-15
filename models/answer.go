package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	CandidatId     uint64  `json:"candidat_id" validate:"required"`
	TestQuestionId uint64  `json:"test_question_id" validate:"required"`
	ChoicesId      uint64  `json:"choices_id"`
	AnswerText     string  `json:"answer_text"`
	AnswerFile     string  `json:"answer_file"`
	Point          float64 `json:"point"`
}
type AnswerDbResponse struct {
	Id       uint64
	Type     string
	Points   int
	IsAnswer bool
	ChoiceId uint64
}
