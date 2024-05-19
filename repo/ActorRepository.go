package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type ActorRepository struct {
	DatabaseConnection *gorm.DB
}

func NewActorRepository(db *gorm.DB) *ActorRepository {
	return &ActorRepository{DatabaseConnection: db}
}

func (ar *ActorRepository) Create(actor *model.Actor) (*model.Actor, error) {
	if err := ar.DatabaseConnection.Create(actor).Error; err != nil {
		return nil, err
	}
	return actor, nil
}

func (ar *ActorRepository) GetAll() ([]model.Actor, error) {
	var actors []model.Actor
	if err := ar.DatabaseConnection.Find(&actors).Error; err != nil {
		return nil, err
	}
	return actors, nil
}
