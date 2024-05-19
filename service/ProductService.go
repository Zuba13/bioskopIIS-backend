package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type ProductService struct {
	ProductRepository repo.ProductRepository
}

func NewProductService(productRepository repo.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (productService *ProductService) CreateProduct(contract model.Product) (model.Product, error) {
	return productService.ProductRepository.Create(&contract)
}

func (productService *ProductService) GetAllProducts() ([]model.Product, error) {
	return productService.ProductRepository.GetAll()
}

func (productService *ProductService) GetProductById(id uint) (model.Product, error) {
	return productService.ProductRepository.GetById(id)
}

func (productService *ProductService) UpdateProduct(contract *model.Product) error {
	return productService.ProductRepository.Update(contract)
}

func (productService *ProductService) DeleteProduct(id uint) error {
	return productService.ProductRepository.Delete(id)
}