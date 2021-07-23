package models

type Choices struct {
	Id            uint64          `json:"id" gorm:"primaryKey;autoIncrement" swaggerignore:"true"`
	ChoiceText    *string         `json:"choice_text" validate:"required"`
	QuestionId    uint64          `json:"question_id" swaggerignore:"true"`
	IsAnswer      *bool           `json:"is_answer" validate:"required"`
	AnswerChoices []AnswerChoices `json:"answer_choices" swaggerignore:"true"`
}
