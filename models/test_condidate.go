package models

import (
	"gorm.io/gorm"
	"time"
)

type TestCandidate struct {
	CandidateID     uint64         `json:"candidate_id" validate:"required"gorm:"primaryKey"`
	TestID          uint64         `json:"test_id" validate:"required"gorm:"primaryKey"`
	TestStatus      string         `json:"test_status" gorm:"default:waiting"`
	Score           float64        `json:"score"`
	TimeLimit       uint64         `json:"time_limit"`
	CurrentQuestion uint64         `json:"current_question"`
	Answer          []Answer       `gorm:"ForeignKey:TestID,CandidateID;References:TestID,CandidateID"`
	CreatedAt       time.Time      `swaggerignore:"true"`
	UpdatedAt       time.Time      `swaggerignore:"true"`
	DeletedAt       gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}
type TestsCandidatesResponse struct {
	TestCandidateID string  `json:"test_candidate_id"`
	TestName        string  `json:"test_name"`
	CandidateName   string  `json:"candidate_name"`
	CandidateEmail  string  `json:"candidate_email"`
	Score           float64 `json:"score"`
	TestStatus      string  `json:"test_status"`
}
type StartTest struct {
	Name            string               `json:"name"`
	Email           string               `json:"email"`
	Questions       []StartTestQuestions `json:"questions"`
	TestStatus      string               `json:"test_status"`
	Score           float64              `json:"score"`
	CurrentQuestion uint64               `json:"current_question"`
}
type StartTestQuestions struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	ExpectedTime *int   `json:"expected_time"`
}
type Quiz struct {
	Test
}
type UpdateTestStatus struct {
	TestStatus string `json:"test_status"`
}
type UpdateCurrentQuestion struct {
	CurrentQuestion uint64 `json:"current_question"`
}
