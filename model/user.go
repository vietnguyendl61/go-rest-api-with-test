package model

import (
	"time"
)

type User struct {
	BaseModel
	BookDate    time.Time `json:"book_date" gorm:"column:book_date;not null"`
	Description string    `json:"description" gorm:"column:description;type:text"`
}

func (u User) TableName() string {
	return "users"
}
