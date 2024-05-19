package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type ContractItemRepository struct {
	DatabaseConnection *gorm.DB
}

func NewContractItemRepository(db *gorm.DB) *ContractItemRepository {
	return &ContractItemRepository{DatabaseConnection: db}
}

func (cr *ContractItemRepository) Create(contractItem *model.ContractItem) (model.ContractItem, error) {
	if err := cr.DatabaseConnection.Create(contractItem).Error; err != nil {
		return model.ContractItem{}, err
	}
	return *contractItem, nil
}

func (cr *ContractItemRepository) GetAll() ([]model.ContractItem, error) {
	var contractItems []model.ContractItem
	if err := cr.DatabaseConnection.Find(&contractItems).Error; err != nil {
		return nil, err
	}
	return contractItems, nil
}

func (cr *ContractItemRepository) GetByID(id uint) (model.ContractItem, error) {
	var contractItem model.ContractItem
	if err := cr.DatabaseConnection.First(&contractItem, id).Error; err != nil {
		return model.ContractItem{}, err
	}
	return contractItem, nil
}

func (cr *ContractItemRepository) Update(contractItem *model.ContractItem) error {
	return cr.DatabaseConnection.Save(contractItem).Error
}

func (cr *ContractItemRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.ContractItem{}, id).Error
}