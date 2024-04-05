package repo

import (
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

func (mr *MovieRepository) GetByID(id uint) (model.Movie, error) {
	var movie model.Movie
	if err := mr.DatabaseConnection.First(&movie, id).Error; err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func (mr *MovieRepository) Update(movie *model.Movie) error {
	return mr.DatabaseConnection.Save(movie).Error
}

func (mr *MovieRepository) Delete(id uint) error {
	return mr.DatabaseConnection.Delete(&model.Movie{}, id).Error
}
