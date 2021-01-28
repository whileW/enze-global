package utils

import "time"

type BaseModel struct {
	ID 			int			`json:"id" gorm:"primary_key"`
	CreatedAt time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time	`json:"updated_at" gorm:"type:datetime"`
	DeletedAt *time.Time	`json:"deleted_at" sql:"index"`
}

type BaseModelV1 struct {
	CreatedAt time.Time	`json:"created_at" gorm:"type:datetime"`
	UpdatedAt time.Time	`json:"updated_at" gorm:"type:datetime"`
	DeletedAt *time.Time	`json:"deleted_at" sql:"index"`
}


