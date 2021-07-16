package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	CandidateId *uint64 `json:"candidate_id"`
	UserId      *uint64 `json:"user_id" validate:"required"`
	Comment     string  `json:"comment" validate:"required"`
}
