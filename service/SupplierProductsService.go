package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type SupplierProductService struct {
	SupplierProductRepository repo.SupplierProductRepository
}

func NewSupplierProductService(supplierProductRepository repo.SupplierProductRepository) *SupplierProductService {
	return &SupplierProductService{
		SupplierProductRepository: supplierProductRepository,
	}
}

func (supplierProductService *SupplierProductService) CreateSupplierProduct(contract model.SupplierProduct) (model.SupplierProduct, error) {
	return supplierProductService.SupplierProductRepository.Create(&contract)
}

func (supplierProductService *SupplierProductService) GetAllSupplierProducts() ([]model.SupplierProduct, error) {
	return supplierProductService.SupplierProductRepository.GetAll()
}

func (supplierProductService *SupplierProductService) GetSupplierProductById(id uint) (model.SupplierProduct, error) {
	return supplierProductService.SupplierProductRepository.GetById(id)
}

func (supplierProductService *SupplierProductService) UpdateSupplierProduct(contract *model.SupplierProduct) error {
	return supplierProductService.SupplierProductRepository.Update(contract)
}

func (supplierProductService *SupplierProductService) DeleteSupplierProduct(id uint) error {
	return supplierProductService.SupplierProductRepository.Delete(id)
}