package repo

import (
	"time"

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

func (dcr *DistributionContractRepository) IsDateInContract(movieId uint, date time.Time) (bool, error) {
	// Extract the year, month, and day from date and create a new time.Time value with these components and a time of 00:00:00
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	var count int64
	err := dcr.DatabaseConnection.
		Model(&model.DistributionContract{}).
		Where("movie_id = ? AND date(start_date) <= ? AND date(end_date) >= ?", movieId, date, date).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
