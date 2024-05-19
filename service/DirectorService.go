package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type DirectorService struct {
	DirectorRepository repo.DirectorRepository
}

func NewDirectorService(directorRepository repo.DirectorRepository) *DirectorService {
	return &DirectorService{
		DirectorRepository: directorRepository,
	}
}

func (ds *DirectorService) CreateDirector(director *model.Director) (*model.Director, error) {
	return ds.DirectorRepository.Create(director)
}

func (ds *DirectorService) GetAllDirectors() ([]model.Director, error) {
	return ds.DirectorRepository.GetAll()
}
