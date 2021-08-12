package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	CandidateID   uint64          `json:"candidate_id"gorm:"ForeignKey:CandidateID;References:CandidateID"`
	TestID        uint64          `json:"test_id"gorm:"ForeignKey:TestID;References:TestID"`
	QuestionID    uint64          `json:"question_id" validate:"required"`
	AnswerText    string          `json:"answer_text"`
	AnswerFile    string          `json:"answer_file"`
	Point         float64         `json:"point"`
	AnswerChoices []AnswerChoices `json:"answer_choices"`
}
