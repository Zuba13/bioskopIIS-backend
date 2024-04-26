package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type ContractService struct {
	ContractRepository repo.ContractRepository
}

func NewContractService(contractRepository repo.ContractRepository) *ContractService {
	return &ContractService{
		ContractRepository: contractRepository,
	}
}

func (contractService *ContractService) CreateContract(contract model.Contract) (model.Contract, error) {
	return contractService.ContractRepository.Create(&contract)
}

func (contractService *ContractService) GetAllContracts() ([]model.Contract, error) {
	return contractService.ContractRepository.GetAll()
}

func (contractService *ContractService) GetContractById(id uint) (model.Contract, error) {
	return contractService.ContractRepository.GetByID(id)
}

func (contractService *ContractService) UpdateContract(contract *model.Contract) error {
	return contractService.ContractRepository.Update(contract)
}

func (contractService *ContractService) DeleteContract(id uint) error {
	return contractService.ContractRepository.Delete(id)
}