package repo

import (
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

func (pr *ProjectionRepository) UpdateProjection(projection *model.Projection) error {
	return pr.DatabaseConnection.Save(projection).Error
}

func (pr *ProjectionRepository) DeleteProjection(id uint) error {
	return pr.DatabaseConnection.Delete(&model.Projection{}, id).Error
}
