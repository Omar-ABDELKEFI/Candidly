package models

import "gorm.io/gorm"

type Test struct {
	gorm.Model
	Name         string         `json:"name" validate:"required" gorm:"unique"`
	Description  string         `json:"description" validate:"required"`
	TimingPolicy string         `json:"timingPolicy" validate:"required,oneof=Normal Strict Relaxed"`
	ShowScore    *bool          `json:"showScore" validate:"required"`
	Archived     *bool          `json:"archived" validate:"required"`
	PassingScore *uint16        `json:"passingScore" validate:"required"`
	NotifyEmails []NotifyEmails `json:"notifyEmails"`
}
type NotifyEmails struct {
	gorm.Model
	TestId uint64 `json:"testId"`
	Email  string `json:"email" validate:"required,email"`
}
