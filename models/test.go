package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name          string          `json:"name" gorm:"unique"`
	Description   string          `json:"description"`
	TimingPolicy  string          `json:"timing_policy"`
	ShowScore     *bool           `json:"show_score"`
	Archived      *bool           `json:"archived"`
	PassingScore  *uint16         `json:"passing_score"`
	NotifyEmails  string          `json:"notify_emails"`
	TestQuestion  []TestQuestion  `json:"test_questions" swaggerignore:"true"`
	TestCandidate []TestCandidate `json:"test_candidate" swaggerignore:"true"`
}
type CreateTestInput struct {
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	TimingPolicy string  `json:"timing_policy" validate:"required,oneof=Normal Strict Relaxed"`
	ShowScore    *bool   `json:"show_score" validate:"required"`
	Archived     *bool   `json:"archived"`
	PassingScore *uint16 `json:"passing_score" validate:"required"`
	NotifyEmails string  `json:"notify_emails"`
}

func NewTest(input CreateTestInput) Test {
	return Test{
		Name:         input.Name,
		Description:  input.Description,
		TimingPolicy: input.TimingPolicy,
		ShowScore:    input.ShowScore,
		Archived:     input.Archived,
		PassingScore: input.PassingScore,
		NotifyEmails: input.NotifyEmails,
	}
}
