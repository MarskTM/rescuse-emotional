package model

import (
	"time"
)


type Message struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	UserName int    `json:"userId" gorm:"column:user_id;unique"`
	Message  string `json:"message"`

	CreatedAt time.Time `swaggerignore:"true"`
}
