package models

type Candidate struct {
	OwnModel
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Test  []*Test `gorm:"many2many:test_candidates;"`
}

type CandidateRequest struct {
	Name  string          `json:"name"`
	Email string          `json:"email" `
	Test  []*TestResponse `gorm:"many2many:test_candidates;"`
}
type CandidateResponse struct {
	OwnModel
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Test  []*Test `gorm:"many2many:test_candidates;"`
}
