package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type SupplierRepository struct {
	DatabaseConnection *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{DatabaseConnection: db}
}

func (sr *SupplierRepository) Create(supplier *model.Supplier) (model.Supplier, error) {
	if err := sr.DatabaseConnection.Create(supplier).Error; err != nil {
		return model.Supplier{}, err
	}
	return *supplier, nil
}

func (sr *SupplierRepository) GetAll() ([]model.Supplier, error) {
	var movies []model.Supplier
	if err := sr.DatabaseConnection.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (sr *SupplierRepository) GetByID(id uint) (model.Supplier, error) {
	var supplier model.Supplier
	if err := sr.DatabaseConnection.First(&supplier, id).Error; err != nil {
		return model.Supplier{}, err
	}
	return supplier, nil
}

func (sr *SupplierRepository) Update(supplier *model.Supplier) error {
	return sr.DatabaseConnection.Save(supplier).Error
}

func (sr *SupplierRepository) Delete(id uint) error {
	return sr.DatabaseConnection.Delete(&model.Supplier{}, id).Error
}