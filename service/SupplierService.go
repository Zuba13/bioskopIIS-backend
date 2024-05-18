package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type SupplierService struct {
	SupplierRepository repo.SupplierRepository
}

func NewSupplierService(movieRepository repo.SupplierRepository) *SupplierService {
	return &SupplierService{
		SupplierRepository: movieRepository,
	}
}

func (supplierService *SupplierService) CreateSupplier(supplier model.Supplier) (model.Supplier, error) {
	return supplierService.SupplierRepository.Create(&supplier)
}

func (supplierService *SupplierService) GetAllSuppliers() ([]model.Supplier, error) {
	return supplierService.SupplierRepository.GetAll()
}

func (supplierService *SupplierService) GetSupplierById(id uint) (model.Supplier, error) {
	return supplierService.SupplierRepository.GetByID(id)
}

func (supplierService *SupplierService) UpdateSupplier(supplier *model.Supplier) error {
	return supplierService.SupplierRepository.Update(supplier)
}

func (supplierService *SupplierService) DeleteSupplier(id uint) error {
	return supplierService.SupplierRepository.Delete(id)
}