package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type TheatreInfoRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *TheatreInfoRepository) GetTheatreInfo() (model.TheatreInfo, error) {
	var info model.TheatreInfo
	if err := repo.DatabaseConnection.First(&info).Error; err != nil {
		return model.TheatreInfo{}, err
	}
	return info, nil
}

func (repo *TheatreInfoRepository) UpdateTheatreInfo(info model.TheatreInfo) error {
	return repo.DatabaseConnection.Save(&info).Error
}
