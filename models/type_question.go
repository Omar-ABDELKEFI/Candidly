package models

import "gorm.io/gorm"

type Choices struct {
	Id         uint64 `json:"id" gorm:"primaryKey;autoIncrement"`
	ChoiceText string `json:"choise_text"`
	McqId      uint64 `json:"mcq_id"`
	IsAnswer   bool   `json:"is_answer"`
}
type Num struct {
	gorm.Model
	id                uint64 `json:"num_id" gorm:"primaryKey;autoIncrement"`
	Default           int    `json:"default" gorm:"column:default"`
	Hint              int    `json:"hint" gorm:"column:hint"`
	InputPrecision    int    `json:"input_precision" gorm:"column:input_precision"`
	Max               int    `json:"max" gorm:"column:max"`
	Min               int    `json:"min" gorm:"column:min"`
	Prefix            string `json:"prefix" gorm:"column:prefix"`
	Suffix            string `json:"Suffix" gorm:"column:Suffix"`
	RoundingPrecision int    `json:"RoundingPrecon" gorm:"column:rounding_precision"`
	Solution          int    `json:"default" gorm:"column:solution"`
	QuestionId        uint64 `json:"question_id"`
}
type MultipleChoiceQuestion struct {
	Id         uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	QuestionId uint64    `json:"question_id"`
	Choices    []Choices `gorm:"foreignkey:McqId;references:id"`
}

type Text struct {
	id             uint64 `json:"text_id" gorm:"primaryKey;autoIncrement"`
	QuestionId     uint64 `json:"question_id"`
	AnswerTemplate string `json:"answer_template"`
}
