package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type ContractItemService struct {
	ContractItemRepository repo.ContractItemRepository
}

func NewContractItemItemService(contractItemRepository repo.ContractItemRepository) *ContractItemService {
	return &ContractItemService{
		ContractItemRepository: contractItemRepository,
	}
}

func (contractItemService *ContractItemService) CreateContractItem(contract model.ContractItem) (model.ContractItem, error) {
	return contractItemService.ContractItemRepository.Create(&contract)
}

func (contractItemService *ContractItemService) GetAllContractItems() ([]model.ContractItem, error) {
	return contractItemService.ContractItemRepository.GetAll()
}

func (contractItemService *ContractItemService) GetContractItemById(id uint) (model.ContractItem, error) {
	return contractItemService.ContractItemRepository.GetByID(id)
}

func (contractItemService *ContractItemService) UpdateContractItem(contract *model.ContractItem) error {
	return contractItemService.ContractItemRepository.Update(contract)
}

func (contractItemService *ContractItemService) DeleteContractItem(id uint) error {
	return contractItemService.ContractItemRepository.Delete(id)
}