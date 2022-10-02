package service

import (
	"rescues/model"
	"rescues/repository"
)

type AnswerService interface {
	GetById(id int) (*model.Answer, error)
	GetAll() ([]model.Answer, error)
	Create(newAnswer *model.Answer) (*model.Answer, error)
	Update(newAnswer model.Answer) (*model.Answer, error)
	Delete(id int) (error)
	FilterByQuestionGroup(groupId int) ([]model.Answer, error)
}

type answerService struct {
	answerRepo model.AnswerRepository
}

// --------------------answer module------------------------

func (s *answerService) GetById(id int) (*model.Answer, error) {
	answer, err := s.answerRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (s *answerService) GetAll() ([]model.Answer, error) {
	answers, err := s.answerRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func (s *answerService) Create(new *model.Answer) (*model.Answer, error) {
	answer, err := s.answerRepo.Create(new)
	if err != nil {
		return nil, err
	}
	return answer, nil
}

func (s *answerService) Update(answer model.Answer) (*model.Answer, error) {
	record, err := s.answerRepo.Update(answer)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (s *answerService) Delete(id int) error {
	if err := s.answerRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

func (s *answerService) FilterByQuestionGroup(groupId int) ([]model.Answer, error) {
	answers, err := s.answerRepo.FilterByQuestionGroup(groupId)
	if err != nil {
		return nil, err
	}
	return answers, nil
}

func NewAnswerService() *answerService {
	answerRepo := repository.NewAnswerRepository()
	return &answerService{
		answerRepo: answerRepo,
	}
}
