package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type HallService struct {
	HallRepository repo.HallRepository
}

func NewHallService(hallRepository repo.HallRepository) *HallService {
	return &HallService{
		HallRepository: hallRepository,
	}
}

func (hs *HallService) CreateHall(hall model.Hall) (model.Hall, error) {
	return hs.HallRepository.Create(&hall)
}

func (hs *HallService) GetAllHalls() ([]model.Hall, error) {
	return hs.HallRepository.GetAll()
}

func (hs *HallService) GetHallByID(id uint) (model.Hall, error) {
	return hs.HallRepository.GetByID(id)
}

func (hs *HallService) UpdateHall(hall *model.Hall) error {
	return hs.HallRepository.Update(hall)
}

func (hs *HallService) DeleteHall(id uint) error {
	return hs.HallRepository.Delete(id)
}
