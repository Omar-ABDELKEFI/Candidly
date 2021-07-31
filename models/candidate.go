package models

type Candidate struct {
	OwnModel
	Name  string  `json:"name"`
	Email string  `json:"email" gorm:"unique"`
	Test  []*Test `gorm:"many2many:test_candidates;"`
}
