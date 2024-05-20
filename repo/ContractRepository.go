package repo

import (
	"time"

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
	var contracts []model.Contract
	if err := cr.DatabaseConnection.Preload("ContractItems.Product").Preload("Supplier").Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func (cr *ContractRepository) GetAllSupplierContracts(supplierId uint) ([]model.Contract, error) {
	var contracts []model.Contract
	if err := cr.DatabaseConnection.Preload("ContractItems.Product").Preload("Supplier").Where("supplier_id = ?", supplierId).Find(&contracts).Error; err != nil {
		return nil, err
	}
	return contracts, nil
}

func (cr *ContractRepository) GetByID(id uint) (model.Contract, error) {
	var contract model.Contract
	if err := cr.DatabaseConnection.Preload("ContractItems.Product").Preload("Supplier").First(&contract, id).Error; err != nil {
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

func (cr *ContractRepository) GetAllContractForTodayDelivery() ([]model.Contract, error) {
	today := time.Now()
    startOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, today.Location())
    endOfDay := startOfDay.Add(24 * time.Hour).Add(-time.Nanosecond) 
    var atOnceContracts []model.Contract
    // Query contracts where valid_from falls within today and contract_type is 1
    if err := cr.DatabaseConnection.
        Preload("ContractItems.Product").
        Where("valid_from BETWEEN ? AND ? AND contract_type = ?", startOfDay, endOfDay, 1).
        Find(&atOnceContracts).Error; err != nil {
        return nil, err
    }
	return atOnceContracts,nil
}

func (cr *ContractRepository) GetAllContractFromWeeklyDelivery() ([]model.Contract, error) {
	var weeklyContracts []model.Contract
    if err := cr.DatabaseConnection.
        Preload("ContractItems.Product").
        Where("contract_type = ? AND EXTRACT(ISODOW FROM  valid_from) = EXTRACT(ISODOW FROM CURRENT_DATE)", 0).
        Find(&weeklyContracts).Error; err != nil {
        return nil, err
    }
    return weeklyContracts, nil
}