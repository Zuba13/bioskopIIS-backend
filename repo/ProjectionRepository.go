package repo

import (
	"time"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type ProjectionRepository struct {
	DatabaseConnection *gorm.DB
}

func NewProjectionRepository(db *gorm.DB) *ProjectionRepository {
	return &ProjectionRepository{DatabaseConnection: db}
}

func (pr *ProjectionRepository) CreateProjection(projection model.Projection) (*model.Projection, error) {
	if err := pr.DatabaseConnection.Create(&projection).Error; err != nil {
		return nil, err
	}
	return &projection, nil
}

func (pr *ProjectionRepository) GetAllProjections() ([]model.Projection, error) {
	var projections []model.Projection
	if err := pr.DatabaseConnection.Find(&projections).Error; err != nil {
		return nil, err
	}
	return projections, nil
}

func (pr *ProjectionRepository) GetProjectionByID(id uint) (*model.Projection, error) {
	var projection model.Projection
	if err := pr.DatabaseConnection.First(&projection, id).Error; err != nil {
		return nil, err
	}
	return &projection, nil
}

func (pr *ProjectionRepository) UpdateProjection(projection *model.Projection, db ...*gorm.DB) (*model.Projection, error) {
	var err error
	if len(db) > 0 {
		err = db[0].Save(projection).Error
	} else {
		err = pr.DatabaseConnection.Save(projection).Error
	}
	if err != nil {
		return nil, err
	}
	return projection, nil
}

func (pr *ProjectionRepository) DeleteProjection(id uint) error {
	return pr.DatabaseConnection.Delete(&model.Projection{}, id).Error
}

// Implementation specific for PostgreSQL
func (pr *ProjectionRepository) GetFilteredProjections(hallId, movieId uint, date time.Time) ([]model.Projection, error) {
	var projections []model.Projection
	db := pr.DatabaseConnection

	if hallId != 0 {
		db = db.Where("hall_id = ?", hallId)
	}

	if !date.IsZero() {
		db = db.Where("date(timeslot_start_time) = ?", date.Format("2006-01-02"))
	}

	if movieId != 0 {
		db = db.Where("movie_id = ?", movieId)
	}

	db = db.Where("is_canceled = ?", false)

	if err := db.Find(&projections).Error; err != nil {
		return nil, err
	}

	return projections, nil
}

func (pr *ProjectionRepository) GetProjectionByDateAndIds(startTime time.Time, endTime time.Time, movieId uint, hallId uint) (*model.Projection, error) {
	var projection model.Projection
	err := pr.DatabaseConnection.
		Model(&model.Projection{}).
		Where("movie_id = ? AND hall_id = ? AND timeslot_start_time = ? AND timeslot_end_time = ? AND is_canceled = false",
			movieId, hallId, startTime, endTime).
		First(&projection).
		Error

	if err != nil {
		return nil, err
	}

	return &projection, nil
}
