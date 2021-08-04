package models

type Test struct {
	OwnModel
	Name         string         `json:"name"`
	Description  string         `json:"description"`
	TimingPolicy string         `json:"timing_policy"`
	ShowScore    *bool          `json:"show_score"`
	Archived     *bool          `json:"archived"`
	PassingScore *uint16        `json:"passing_score"`
	NotifyEmails string         `json:"notify_emails"`
	TestQuestion []TestQuestion `json:"test_questions"`
	Candidate    []*Candidate   `gorm:"many2many:test_candidate;"`
}
type TestRequest struct {
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	TimingPolicy string  `json:"timing_policy"`
	ShowScore    *bool   `json:"show_score"`
	Archived     *bool   `json:"archived"`
	PassingScore *uint16 `json:"passing_score"`
	NotifyEmails string  `json:"notify_emails"`
}
type TestResponse struct {
	OwnModel
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	TimingPolicy string  `json:"timing_policy"`
	ShowScore    *bool   `json:"show_score"`
	Archived     *bool   `json:"archived"`
	PassingScore *uint16 `json:"passing_score"`
	NotifyEmails string  `json:"notify_emails"`
}
