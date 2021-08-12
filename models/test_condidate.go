package models

import (
	"gorm.io/gorm"
	"time"
)

type TestCandidate struct {
	CandidateID uint64         `json:"candidate_id" validate:"required"gorm:"primaryKey"`
	TestID      uint64         `json:"test_id" validate:"required"gorm:"primaryKey"`
	TestStatus  string         `json:"test_status" gorm:"default:waiting"`
	Score       float64        `json:"score"`
	Answer      []Answer       `gorm:"ForeignKey:TestID,CandidateID;References:TestID,CandidateID"`
	CreatedAt   time.Time      `swaggerignore:"true"`
	UpdatedAt   time.Time      `swaggerignore:"true"`
	DeletedAt   gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}
type TestsCandidatesResponse struct {
	TestName       string
	CandidateName  string
	CandidateEmail string
	Score          float64
	TestStatus     string
}
type Quiz struct {
	Test
}
