package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	CandidateTestId uint64          `json:"candidate_id" validate:"required"`
	QuestionId      uint64          `json:"question_id" validate:"required"`
	AnswerText      string          `json:"answer_text"`
	AnswerFile      string          `json:"answer_file"`
	Point           float64         `json:"point"`
	AnswerChoices   []AnswerChoices `json:"answer_choices"`
}
