package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	CandidateId    uint64          `json:"candidate_id" validate:"required"`
	TestQuestionId uint64          `json:"test_question_id" validate:"required"`
	AnswerText     string          `json:"answer_text"`
	AnswerFile     string          `json:"answer_file"`
	Point          float64         `json:"point"`
	AnswerChoices  []AnswerChoices `json:"answer_choices"`
}
type QuestionChoices struct {
	QuestionId uint64
	Type       string
	Points     int
	IsAnswer   bool
	ChoiceId   uint64
}
