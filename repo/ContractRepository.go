package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type ContractRepository struct {
	DatabaseConnection *gorm.DB
}

func NewContractRepository(db *gorm.DB) *ContractRepository {
	return &ContractRepository{DatabaseConnection: db}
}

func (cr *ContractRepository) Create(contract *model.Contract) (model.Contract, error) {
	if err := cr.DatabaseConnection.Create(contract).Error; err != nil {
		return model.Contract{}, err
	}
	return *contract, nil
}

func (cr *ContractRepository) GetAll() ([]model.Contract, error) {
	var movies []model.Contract
	if err := cr.DatabaseConnection.Preload("ContractItems").Preload("Supplier").Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (cr *ContractRepository) GetByID(id uint) (model.Contract, error) {
	var contract model.Contract
	if err := cr.DatabaseConnection.Preload("ContractItems").Preload("Supplier").First(&contract, id).Error; err != nil {
		return model.Contract{}, err
	}
	return contract, nil
}

func (cr *ContractRepository) Update(contract *model.Contract) error {
	return cr.DatabaseConnection.Save(contract).Error
}

func (cr *ContractRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.Contract{}, id).Error
}