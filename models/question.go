package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique" validate:"required,min=1"`
	Question Question
}

type Question struct {
	gorm.Model
	Id                     uint64                 `json:"id" gorm:"column=id;primaryKey;autoIncrement"`
	Difficulty             string                 `json:"difficulty" validate:"required,oneof=hard eazy"`
	Points                 int                    `json:"points" validate:"required,min=1,max=6"`
	Name                   string                 `json:"name" validate:"required" gorm:"unique"`
	SkillId                uint64                 `json:"skill_id"`
	ExpectedTime           int                    `json:"expected_time" validate:"required,oneof=1 2 3 5 7 10 15 20 30 40 60"`
	QuestionText           string                 `json:"question_text" valiate:"required" `
	Num                    Num                    `gorm:"foreignkey:QuestionId;references:id"`
	MultipleChoiceQuestion MultipleChoiceQuestion `gorm:"foreignkey:QuestionId;references:id"`
	Text                   Text                   `gorm:"foreignkey:QuestionId;references:id"`
}
