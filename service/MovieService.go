package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type MovieService struct {
	MovieRepository repo.MovieRepository
}

func NewMovieService(movieRepository repo.MovieRepository) *MovieService {
	return &MovieService{
		MovieRepository: movieRepository,
	}
}

func (ms *MovieService) CreateMovie(movie model.Movie) (model.Movie, error) {
	return ms.MovieRepository.Create(&movie)
}

func (ms *MovieService) GetAllMovies() ([]model.Movie, error) {
	return ms.MovieRepository.GetAll()
}

func (ms *MovieService) GetMovieByID(id uint) (model.Movie, error) {
	return ms.MovieRepository.GetByID(id)
}

func (ms *MovieService) UpdateMovie(movie *model.Movie) error {
	return ms.MovieRepository.Update(movie)
}

func (ms *MovieService) DeleteMovie(id uint) error {
	return ms.MovieRepository.Delete(id)
}

func (ms *MovieService) GetProjectionsForMovie(movieID uint) ([]model.Projection, error) {
	return ms.MovieRepository.GetProjectionsForMovie(movieID)
}
