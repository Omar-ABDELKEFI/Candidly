package models

type Skill struct {
	OwnModel
	Name string `json:"name"`
}
type CreateSkillInput struct {
	Name string `json:"name" validate:"required,min=1"`
}

func NewSkill(input CreateSkillInput) Skill {
	return Skill{
		Name: input.Name,
	}
}

type SkillsResponse struct {
	ID   uint64 `json:"ID" validate:"required"`
	Name string `json:"name" validate:"required"`
}
