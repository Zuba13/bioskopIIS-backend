package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type SupplierProductRepository struct {
	DatabaseConnection *gorm.DB
}

func NewSupplierProductRepository(db *gorm.DB) *SupplierProductRepository {
	return &SupplierProductRepository{DatabaseConnection: db}
}

func (cr *SupplierProductRepository) Create(supplierProduct *model.SupplierProduct) (model.SupplierProduct, error) {
	if err := cr.DatabaseConnection.Create(supplierProduct).Error; err != nil {
		return model.SupplierProduct{}, err
	}
	return *supplierProduct, nil
}

func (cr *SupplierProductRepository) GetAll() ([]model.SupplierProduct, error) {
	var supplierProducts []model.SupplierProduct
	if err := cr.DatabaseConnection.Find(&supplierProducts).Error; err != nil {
		return nil, err
	}
	return supplierProducts, nil
}

func (cr *SupplierProductRepository) GetById(id uint) (model.SupplierProduct, error) {
	var supplierProduct model.SupplierProduct
	if err := cr.DatabaseConnection.First(&supplierProduct, id).Error; err != nil {
		return model.SupplierProduct{}, err
	}
	return supplierProduct, nil
}

func (cr *SupplierProductRepository) Update(supplierProduct *model.SupplierProduct) error {
	return cr.DatabaseConnection.Save(supplierProduct).Error
}

func (cr *SupplierProductRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.SupplierProduct{}, id).Error
}

func (cr *SupplierProductRepository) GetByProductId(productId uint) (model.SupplierProduct, error) {
	var supplierProduct model.SupplierProduct
	if err := cr.DatabaseConnection.First(&supplierProduct, productId).Error; err != nil {
		return model.SupplierProduct{}, err
	}
	return supplierProduct, nil
}