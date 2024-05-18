package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type DirectorRepository struct {
	DatabaseConnection *gorm.DB
}

func NewDirectorRepository(db *gorm.DB) *DirectorRepository {
	return &DirectorRepository{DatabaseConnection: db}
}

func (dr *DirectorRepository) Create(director *model.Director) (*model.Director, error) {
	if err := dr.DatabaseConnection.Create(director).Error; err != nil {
		return nil, err
	}
	return director, nil
}

func (dr *DirectorRepository) GetAll() ([]model.Director, error) {
	var directors []model.Director
	if err := dr.DatabaseConnection.Find(&directors).Error; err != nil {
		return nil, err
	}
	return directors, nil
}
