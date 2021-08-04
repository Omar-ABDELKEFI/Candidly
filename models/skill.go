package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name string `json:"name"`
}
type CreateSkillInput struct {
	Name string `json:"name" validate:"required,min=1"`
}

func NewSkill(input CreateSkillInput) Skill {
	return Skill{
		Name: input.Name,
	}
}
