package models

type Test struct {
	OwnModel
	Name         string       `json:"name"`
	Description  string       `json:"description"`
	TimeLimit    uint64       `json:"time_limit"`
	TimingPolicy string       `json:"timing_policy"`
	ShowScore    *bool        `json:"show_score"`
	Archived     *bool        `json:"archived"`
	PassingScore *uint16      `json:"passing_score"`
	NotifyEmails string       `json:"notify_emails"`
	Questions    []*Question  `json:"questions" gorm:"many2many:test_questions;"`
	Candidate    []*Candidate `json:"candidate" gorm:"many2many:test_candidates;"`
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
type MyTests struct {
	TestId          *int    `json:"test_id"`
	NumberQuestion  *int    `json:"number_Question"`
	NumberCandidate *int    `json:"number_candidate"`
	TestName        *string `json:"test_name"`
	ExpectedTime    *int    `json:"expected_time"`
}
