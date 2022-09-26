package repository

import (
	"time"
	"rescues/model"
	"rescues/infrastructure"
)

type answerRepository struct {}

func (r *answerRepository) GetById(id int) (*model.Answer, error) {
	db := infrastructure.GetDB()
	var answer model.Answer
	if err := db.Where("id = ?", id).First(&answer).Error; err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *answerRepository) GetAll() ([]model.Answer, error) {
	db := infrastructure.GetDB()
	var answers []model.Answer

	if err := db.Model(&model.Judge{}).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

func (r *answerRepository) Create(newAnswer *model.Answer) (*model.Answer, error) {
	db := infrastructure.GetDB()
	if err := db.Create(newAnswer).Error; err != nil {
		return nil, err
	}
	return newAnswer, nil
}

func (r *answerRepository) Update(newAnswer model.Answer) (*model.Answer, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&newAnswer).Where("id = ?", newAnswer.Id).Updates(newAnswer).Error; err != nil {
		return nil, err
	}

	var answer model.Answer
	if err := db.Where("id = ?", newAnswer.Id).First(&answer).Error; err != nil {
		return nil, err
	}
	return &answer, nil
}

func (r *answerRepository) Delete(id int) (error) {
	db := infrastructure.GetDB()

	if err := db.Model(&model.Answer{Id: id}).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}

func (r *answerRepository) FilterByQuestionId(questionId int) ([]model.Answer, error) {
	db := infrastructure.GetDB()
	var answers []model.Answer

	if err := db.Where("question_id = ?", questionId).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil

}

func NewAnswerRepository() *answerRepository {
	return &answerRepository{}
}
