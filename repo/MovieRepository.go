package repo

import (
	"strings"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type MovieRepository struct {
	DatabaseConnection *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{DatabaseConnection: db}
}

func (mr *MovieRepository) Create(movie *model.Movie) (model.Movie, error) {
	if err := mr.DatabaseConnection.Create(movie).Error; err != nil {
		return model.Movie{}, err
	}
	return *movie, nil
}

func (mr *MovieRepository) GetAll() ([]model.Movie, error) {
	var movies []model.Movie
	if err := mr.DatabaseConnection.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (mr *MovieRepository) GetFiltered(title string, year int) ([]model.Movie, error) {
	var movies []model.Movie
	db := mr.DatabaseConnection

	if title != "" {
		db = db.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}
	if year != 0 {
		db = db.Where("year = ?", year)
	}

	if err := db.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (mr *MovieRepository) GetByID(id uint) (model.Movie, error) {
	var movie model.Movie
	if err := mr.DatabaseConnection.First(&movie, id).Error; err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func (mr *MovieRepository) GetByIDWithAssociations(id uint) (model.Movie, error) {
	var movie model.Movie
	if err := mr.DatabaseConnection.Preload("Actors").Preload("Directors").Preload("Contracts.Company").First(&movie, id).Error; err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func (mr *MovieRepository) Update(movie *model.Movie) (*model.Movie, error) {
	db := mr.DatabaseConnection

	if err := db.Save(movie).Error; err != nil {
		return nil, err
	}

	if err := db.Model(movie).Association("Actors").Replace(movie.Actors); err != nil {
		return nil, err
	}

	if err := db.Model(movie).Association("Directors").Replace(movie.Directors); err != nil {
		return nil, err
	}

	updatedMovie, err := mr.GetByIDWithAssociations(movie.ID)
	if err != nil {
		return nil, err
	}

	return &updatedMovie, nil
}

func (mr *MovieRepository) Delete(id uint) error {
	return mr.DatabaseConnection.Delete(&model.Movie{}, id).Error
}

func (mr *MovieRepository) GetProjectionsForMovie(movieID uint) ([]model.Projection, error) {
	var projections []model.Projection
	if err := mr.DatabaseConnection.Where("movie_id = ?", movieID).Find(&projections).Error; err != nil {
		return nil, err
	}
	return projections, nil
}
