package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type Num struct {
	ID                uint64 `gorm:"primaryKey;autoIncrement"`
	Defeult           int
	Hint              int
	InputPrecision    int
	Max               int
	Min               int
	Prefix            string
	Suffix            string
	RoundingPrecision int
	Solutin           int
}
