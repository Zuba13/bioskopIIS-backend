package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type DistributionCompanyRepository struct {
	DatabaseConnection *gorm.DB
}

func NewDistributionCompanyRepository(db *gorm.DB) *DistributionCompanyRepository {
	return &DistributionCompanyRepository{DatabaseConnection: db}
}

func (dcr *DistributionCompanyRepository) Create(company *model.DistributionCompany) (*model.DistributionCompany, error) {
	if err := dcr.DatabaseConnection.Create(company).Error; err != nil {
		return nil, err
	}
	return company, nil
}

func (dcr *DistributionCompanyRepository) GetAll() ([]model.DistributionCompany, error) {
	var companies []model.DistributionCompany
	if err := dcr.DatabaseConnection.Find(&companies).Error; err != nil {
		return nil, err
	}
	return companies, nil
}
