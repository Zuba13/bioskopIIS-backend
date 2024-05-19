package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type CatalogService struct {
	MovieRepository    repo.MovieRepository
	ContractReposiotry repo.DistributionContractRepository
}

func NewCatalogService(movieRepository repo.MovieRepository, contractRepository repo.DistributionContractRepository) *CatalogService {
	return &CatalogService{
		MovieRepository:    movieRepository,
		ContractReposiotry: contractRepository,
	}
}

func (ms *CatalogService) GetFilteredMovies(title string, year int, onlyActive bool) ([]model.Movie, error) {
	movies, err := ms.MovieRepository.GetFiltered(title, year)
	if err != nil {
		return nil, err
	}

	if !onlyActive {
		return movies, nil
	}

	var activeMovies []model.Movie
	for _, movie := range movies {
		contracts, err := ms.ContractReposiotry.GetAll(movie.ID)
		if err != nil {
			return nil, err
		}

		for _, contract := range contracts {
			if contract.IsActive() {
				activeMovies = append(activeMovies, movie)
				break
			}
		}
	}

	return activeMovies, nil
}

func (cs *CatalogService) GetMovieWithAssociations(id uint) (model.Movie, error) {
	return cs.MovieRepository.GetByIDWithAssociations(id)
}
