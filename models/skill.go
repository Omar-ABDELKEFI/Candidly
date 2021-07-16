package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required,min=1"`
	Question []Question
}
