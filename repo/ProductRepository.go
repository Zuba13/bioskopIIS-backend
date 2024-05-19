package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DatabaseConnection *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DatabaseConnection: db}
}

func (cr *ProductRepository) Create(product *model.Product) (model.Product, error) {
	if err := cr.DatabaseConnection.Create(product).Error; err != nil {
		return model.Product{}, err
	}
	return *product, nil
}

func (cr *ProductRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := cr.DatabaseConnection.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (cr *ProductRepository) GetById(id uint) (model.Product, error) {
	var product model.Product
	if err := cr.DatabaseConnection.First(&product, id).Error; err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (cr *ProductRepository) Update(product *model.Product) error {
	return cr.DatabaseConnection.Save(product).Error
}

func (cr *ProductRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.Product{}, id).Error
}