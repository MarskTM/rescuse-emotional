package model

import (
	"time"
)

type Answer struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	QuestionId int    `json:"question_id"`
	Answer     string `json:"answer"`

	ScoreStress  int `json:"score_stress"`
	ScoreDepress int `json:"score_depess"`
	ScoreAnxiety int `json:"score_anxiety"`

	CreatedAt time.Time  `swaggerignore:"true"`
	UpdatedAt time.Time  `swaggerignore:"true"`
	DeletedAt *time.Time `swaggerignore:"true"`
}

type AnswerRepository interface {
	GetById(id int) (*Answer, error)
	GetAll() ([]Answer, error)
	Create(new *Answer) (*Answer, error)
	Update(answer Answer) (*Answer, error)
	Delete(id int) error
	FilterByQuestionId(questionId int) ([]Answer, error)
}
