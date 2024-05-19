package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type DistributionContractRepository struct {
	DatabaseConnection *gorm.DB
}

func NewDistributionContractRepository(db *gorm.DB) *DistributionContractRepository {
	return &DistributionContractRepository{DatabaseConnection: db}
}

func (repo *DistributionContractRepository) Create(contract *model.DistributionContract) (*model.DistributionContract, error) {
	if contract.Company.ID == 0 {
		if err := repo.DatabaseConnection.Create(&contract.Company).Error; err != nil {
			return nil, err
		}
		contract.CompanyID = contract.Company.ID
	}
	err := repo.DatabaseConnection.Create(contract).Error
	if err != nil {
		return nil, err
	}
	return contract, nil
}

func (repo *DistributionContractRepository) Update(contract *model.DistributionContract) (*model.DistributionContract, error) {
	if err := repo.DatabaseConnection.Save(contract).Error; err != nil {
		return nil, err
	}

	if err := repo.DatabaseConnection.First(contract, contract.ID).Error; err != nil {
		return nil, err
	}

	return contract, nil
}

func (repo *DistributionContractRepository) GetAll(movieID uint) ([]model.DistributionContract, error) {
	var contracts []model.DistributionContract
	err := repo.DatabaseConnection.Where("movie_id = ?", movieID).Find(&contracts).Error
	return contracts, err
}

func (repo *DistributionContractRepository) Delete(contract *model.DistributionContract) error {
	return repo.DatabaseConnection.Delete(contract).Error
}

func (repo *DistributionContractRepository) Get(id uint) (*model.DistributionContract, error) {
	var contract model.DistributionContract
	if err := repo.DatabaseConnection.First(&contract, id).Error; err != nil {
		return nil, err
	}
	return &contract, nil
}
