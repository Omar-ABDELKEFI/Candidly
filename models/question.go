package models

type Question struct {
	OwnModel
	Difficulty    string   `json:"difficulty"`
	MaxPoints     *float64 `json:"max_points"`
	Name          string   `json:"name" gorm:"unique"`
	SkillID       uint64   `json:"skill_id"`
	Skill         Skill
	Type          string         `json:"type"`
	ExpectedTime  *int           `json:"expected_time"`
	QuestionText  string         `json:"question_text" `
	FileReadMe    string         `json:"file_read_me"`
	TestQuestions []TestQuestion `json:"test_questions"`
	Choices       []Choices      `json:"choices"`
	Answer        []Answer       `json:"answer"`
}
type CreateQuestionInput struct {
	Difficulty   string    `json:"difficulty" validate:"required,oneof=hard easy"`
	MaxPoints    *float64  `json:"max_points" validate:"required,min=1,max=10"`
	Name         string    `json:"name" validate:"required"`
	SkillID      uint64    `json:"skill_id"`
	Type         string    `json:"type"`
	ExpectedTime *int      `json:"expected_time" validate:"required,oneof=1 2 3 5 7 10 15 20 30 40 60"`
	QuestionText string    `json:"question_text" validate:"required" `
	FileReadMe   string    `json:"file_read_me"`
	Choices      []Choices `json:"choices" validate:"dive"`
	SkillName    string    `json:"skill_name"`
}

func NewQuestion(input CreateQuestionInput) Question {
	return Question{
		Difficulty:   input.Difficulty,
		MaxPoints:    input.MaxPoints,
		Name:         input.Name,
		SkillID:      input.SkillID,
		Type:         input.Type,
		ExpectedTime: input.ExpectedTime,
		QuestionText: input.QuestionText,
		FileReadMe:   input.FileReadMe,
		Choices:      input.Choices,
	}
}
