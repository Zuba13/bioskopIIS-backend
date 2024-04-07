package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type HallRepository struct {
	DatabaseConnection *gorm.DB
}

func NewHallRepository(db *gorm.DB) *HallRepository {
	return &HallRepository{DatabaseConnection: db}
}

func (hr *HallRepository) Create(hall *model.Hall) (model.Hall, error) {
	if err := hr.DatabaseConnection.Create(hall).Error; err != nil {
		return model.Hall{}, err
	}
	return *hall, nil
}

func (hr *HallRepository) GetAll() ([]model.Hall, error) {
	var halls []model.Hall
	if err := hr.DatabaseConnection.Find(&halls).Error; err != nil {
		return nil, err
	}
	return halls, nil
}

func (hr *HallRepository) GetByID(id uint) (model.Hall, error) {
	var hall model.Hall
	if err := hr.DatabaseConnection.First(&hall, id).Error; err != nil {
		return model.Hall{}, err
	}
	return hall, nil
}

func (hr *HallRepository) Update(hall *model.Hall) error {
	return hr.DatabaseConnection.Save(hall).Error
}

func (hr *HallRepository) Delete(id uint) error {
	return hr.DatabaseConnection.Delete(&model.Hall{}, id).Error
}
