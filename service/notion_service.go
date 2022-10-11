package service

import (

	"rescues/model"
	"rescues/repository"
)

type notionService struct{
	notionRepo model.NotionResponse
}

type NotionService interface {
	GetByProfileId(profileId int) ([]model.Notion, error)
	// GetByDate(date time.Time) ([]model.Notion, error)
	Create(notion *model.Notion) (*model.Notion, error)
	Update(notion *model.Notion) (*model.Notion, error)
	Delete(id int) (error)
}

// func (s *notionService) GetByDate(date time.Time) ([]model.Notion, error) {
// 	notions, err := s.notionRepo.GetByDate(date)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return notions, nil
// }

func (s *notionService) GetByProfileId(profileId int) ([]model.Notion, error){
	notions, err := s.notionRepo.GetByProfileId(profileId)
	if err != nil {
		return nil, err
	}
	return notions, nil
}

func (s *notionService) Create(notion *model.Notion) (*model.Notion, error) {
	newNotion, err := s.notionRepo.Create(notion)
	if err != nil {
		return nil, err
	}
	return newNotion, nil
}

func (s *notionService) Update(notion *model.Notion) (*model.Notion, error) {
	newNotion, err := s.notionRepo.Update(notion)
	if err != nil {
		return nil, err
	}
	return newNotion, nil
}

func (s *notionService) Delete(id int) (error) {
	err := s.notionRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func NewNotionService() NotionService {
	notionRepo := repository.NewNotionRepository()
	return &notionService{
		notionRepo: notionRepo,
	}
}
