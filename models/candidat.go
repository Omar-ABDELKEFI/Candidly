package models

import "gorm.io/gorm"

type Candidat struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" validate:"required,email" gorm:"unique"`
	Note  []Note `json:"note"`
}
type Note struct {
	gorm.Model
	CandidatID *uint64 `json:"candidat_id"`
	UserID     *uint64 `json:"user_id" validate:"required"`
	Comment    string  `json:"comment" validate:"required"`
}
