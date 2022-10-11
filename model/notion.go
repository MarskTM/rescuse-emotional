package model

import (
	"time"
)

type Notion struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	ProfileId int       `json:"profileId"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Public    bool      `json:"public"`
	
	CreateAt  time.Time `swaggerignore:"true"`
	UpdatedAt time.Time `swaggerignore:"true"`
}

type NotionResponse interface {
	GetByProfileId(profileId int) ([]Notion, error)
	GetByDate(date time.Time) ([]Notion, error)
	Create(new *Notion) (*Notion, error)
	Update(notion *Notion) (*Notion, error)
	Delete(id int) error
}