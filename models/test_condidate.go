package models

type TestCandidate struct {
	OwnModel
	CandidateId uint64  `json:"candidate_id" validate:"required"`
	TestId      uint64  `json:"test_id" validate:"required"`
	TestStatus  string  `json:"test_status" gorm:"default:waiting"`
	Score       float64 `json:"score"`
}
