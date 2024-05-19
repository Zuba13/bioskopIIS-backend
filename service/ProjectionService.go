package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

// projectionService handles operations related to projections.
type ProjectionService struct {
	ProjectionRepo repo.ProjectionRepository
}

// NewProjectionService creates a new instance of the projectionService.
func NewProjectionService(projectionRepo repo.ProjectionRepository) *ProjectionService {
	return &ProjectionService{
		ProjectionRepo: projectionRepo,
	}
}

// CreateProjection creates a new projection.
func (ps *ProjectionService) CreateProjection(projection model.Projection) (*model.Projection, error) {
	return ps.ProjectionRepo.CreateProjection(projection)
}

// GetAllProjections retrieves all projections.
func (ps *ProjectionService) GetAllProjections() ([]model.Projection, error) {
	return ps.ProjectionRepo.GetAllProjections()
}

// GetProjectionByID retrieves a projection by its ID.
func (ps *ProjectionService) GetProjectionByID(id uint) (*model.Projection, error) {
	return ps.ProjectionRepo.GetProjectionByID(id)
}

// UpdateProjection updates an existing projection.
func (ps *ProjectionService) UpdateProjection(projection *model.Projection) error {
	_, err := ps.ProjectionRepo.UpdateProjection(projection)
	return err
}

// DeleteProjection deletes a projection by its ID.
func (ps *ProjectionService) DeleteProjection(id uint) error {
	return ps.ProjectionRepo.DeleteProjection(id)
}
