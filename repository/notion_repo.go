package repository

import (
	"time"
	"rescues/infrastructure"
	"rescues/model"
)

type notionRepository struct{}

func (r *notionRepository) GetByProfileId(profileId int) ([]model.Notion, error) {
	db := infrastructure.GetDB()
	var notions []model.Notion

	if err := db.Model(&model.Notion{}).Where("profile_id = ?", profileId).Find(&notions).Error; err != nil {
		return nil, err
	}
	return notions, nil
}

func (r *notionRepository) GetByDate(date time.Time) ([]model.Notion, error) {
	db := infrastructure.GetDB()
	
	var notions []model.Notion

	if err := db.Model(&model.Notion{}).Where("create_at = ?", date).Find(&notions).Error; err != nil {
		return nil, err
	}
	return notions, nil
}

func (r *notionRepository) Create(notion *model.Notion) (*model.Notion, error) {
	db := infrastructure.GetDB()

	if err := db.Create(&notion).Error; err != nil {
		return nil, err
	}
	return notion, nil
}

func (r *notionRepository) Update(notion *model.Notion) (*model.Notion, error) {
	db := infrastructure.GetDB()

	if err := db.Model(&notion).Where("id = ?", notion.Id).Updates(notion).Error; err != nil {
		return nil, err
	}
	
	var newNotion model.Notion
	if err := db.Where("id = ?", notion.Id).First(&newNotion).Error; err != nil {
		return nil, err
	}
	return &newNotion, nil
}

func (r *notionRepository) Delete(id int) (error) {
	db := infrastructure.GetDB()

	if err := db.Where("id = ?", id).Delete(&model.Notion{}).Error; err != nil {
		return err
	}
	return nil
}

func NewNotionRepository() *notionRepository {
	return &notionRepository{}
}

