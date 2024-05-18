package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type ActorService struct {
	ActorRepository repo.ActorRepository
}

func NewActorService(actorRepository repo.ActorRepository) *ActorService {
	return &ActorService{
		ActorRepository: actorRepository,
	}
}

func (as *ActorService) CreateActor(actor *model.Actor) (*model.Actor, error) {
	return as.ActorRepository.Create(actor)
}

func (as *ActorService) GetAllActors() ([]model.Actor, error) {
	return as.ActorRepository.GetAll()
}
