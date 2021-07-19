package models

import "gorm.io/gorm"

type Candidate struct {
	gorm.Model
	Name          string          `json:"name"`
	Email         string          `json:"email" gorm:"unique"`
	TestCandidate []TestCandidate `json:"test_candidate" swaggerignore:"true"`
}
type CreateCandidateInput struct {
	Name  string `json:"name"`
	Email string `json:"email" validate:"required,email"`
}

func NewCandidate(input CreateCandidateInput) Candidate {
	return Candidate{
		Name:  input.Name,
		Email: input.Email,
	}
}
