package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name          string          `json:"name"`
	Email         string          `json:"email" validate:"required,email" gorm:"unique"`
	TestCandidate []TestCandidate `json:"test_candidate"`
}
