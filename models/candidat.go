package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name          string          `json:"name"`
	Email         string          `json:"email" validate:"required,email" gorm:"unique"`
	Note          []Note          `json:"note"`
	Answer        []Answer        `json:"answer"`
	TestCandidate []TestCandidate `json:"test_candidate"`
}
type Note struct {
	gorm.Model
	CandidateId uint64 `json:"candidate_id"`
	UserId      uint64 `json:"user_id" validate:"required"`
	Comment     string `json:"comment" validate:"required"`
}
