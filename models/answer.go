package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	CandidatId     uint64 `json:"candidat_id" validate:"required"`
	TestQuestionId uint64 `json:"test_id" validate:"required"`
	Point          int    `json:"point" validate:"required"`
}
