package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type DistributionContractService struct {
	DistributionContractRepository *repo.DistributionContractRepository
	DistributionCompanyRepository  *repo.DistributionCompanyRepository
}

func NewDistributionContractService(distributionContractRepository *repo.DistributionContractRepository, distributionCompanyRepository *repo.DistributionCompanyRepository) *DistributionContractService {
	return &DistributionContractService{
		DistributionContractRepository: distributionContractRepository,
		DistributionCompanyRepository:  distributionCompanyRepository,
	}
}

func (dcs *DistributionContractService) CreateContract(contract *model.DistributionContract, userId uint) (*model.DistributionContract, error) {
	contract.ManagerID = userId
	if err := contract.Validate(); err != nil {
		return nil, &CustomError{Code: ErrBadRequest, Message: err.Error()}
	}
	existingContracts, err := dcs.DistributionContractRepository.GetAll(contract.MovieID)
	if err != nil {
		return nil, &CustomError{Code: ErrInternalServer, Message: err.Error()}
	}
	for _, existingContract := range existingContracts {
		if contract.OverlapsWith(&existingContract) {
			return nil, &CustomError{Code: ErrBadRequest, Message: "contract overlaps with existing contract"}
		}
	}
	return dcs.DistributionContractRepository.Create(contract)
}

func (dcs *DistributionContractService) UpdateContract(contract *model.DistributionContract) (*model.DistributionContract, error) {
	if err := contract.Validate(); err != nil {
		return nil, &CustomError{Code: ErrBadRequest, Message: err.Error()}
	}
	oldContract, err := dcs.DistributionContractRepository.Get(contract.ID)
	if err != nil {
		return nil, &CustomError{Code: ErrInternalServer, Message: err.Error()}
	}
	if oldContract.IsExpired() {
		return nil, &CustomError{Code: ErrForbidden, Message: "expired contracts cannot be updated"}
	}

	return dcs.DistributionContractRepository.Update(contract)
}

func (dcs *DistributionContractService) GetAllCompanies() ([]model.DistributionCompany, error) {
	return dcs.DistributionCompanyRepository.GetAll()
}
