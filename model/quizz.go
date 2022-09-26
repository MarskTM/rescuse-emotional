package model

import (
	"time"
)

type Quizz struct {
	Id            int `json:"id" gorm:"primaryKey"`
	ProfileId     int `json:"profileId"`
	QuestionGroup int `json:"question_group"`

	// Điểm số của bài test
	SumScoreStress  int `json:"sum_score_stress"`
	SumScoreDepress  int `json:"sum_score_depess"`
	SumScoreAnxiety int `json:"sum_score_anxiety"`

	CreatedAt time.Time  `swaggerignore:"true"`
	UpdatedAt time.Time  `swaggerignore:"true"`
	DeletedAt *time.Time `swaggerignore:"true"`
}

type QuizzRepository interface {
	GetById(id int) (*Quizz, error)
	GetAll() ([]Quizz, error)
	Create(new *Quizz) (*Quizz, error)
	Update(quizz Quizz) (*Quizz, error)
	Delete(id int) error
}
