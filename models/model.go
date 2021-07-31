package models

import "time"
import "gorm.io/gorm"

type OwnModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `swaggerignore:"true"`
	UpdatedAt time.Time      `swaggerignore:"true"`
	DeletedAt gorm.DeletedAt `gorm:"index" swaggerignore:"true"`
}
