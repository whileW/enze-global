package utils

import "time"

type BaseModel struct {
	CreatedAt time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time	`json:"updated_at" gorm:"type:datetime"`
	DeletedAt *time.Time	`json:"deleted_at" gorm:"type:datetime"`
}


