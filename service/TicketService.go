package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type TicketService struct {
	TicketRepository *repo.TicketRepository
	UserRepository   *repo.UserRepository
}

func NewTicketService(ticketRepo *repo.TicketRepository, userRepo *repo.UserRepository) *TicketService {
	return &TicketService{
		TicketRepository: ticketRepo,
		UserRepository:   userRepo,
	}
}

func (ts *TicketService) CreateTicket(ticket model.Ticket) (*model.Ticket, error) {
	return ts.TicketRepository.CreateTicket(ticket)
}

func (ts *TicketService) GetAllTickets() ([]*model.Ticket, error) {
	return ts.TicketRepository.GetAllTickets()
}

func (ts *TicketService) GetTicketByID(id uint) (*model.Ticket, error) {
	return ts.TicketRepository.GetTicketByID(id)
}

func (ts *TicketService) UpdateTicket(ticket *model.Ticket) error {
	return ts.TicketRepository.UpdateTicket(ticket)
}

func (ts *TicketService) DeleteTicket(id uint) error {
	return ts.TicketRepository.DeleteTicket(id)
}

func (ts *TicketService) GetTicketsByUserID(userID uint) ([]*model.Ticket, error) {
	return ts.TicketRepository.GetTicketsByUserID(userID)
}

func (ts *TicketService) GetTicketsByProjectionID(projectionId uint) ([]*model.Ticket, error) {
	return ts.TicketRepository.GetTicketsByProjectionID(projectionId)
}

func (ts *TicketService) RefundTickets(projection model.Projection) error {
	tickets, err := ts.TicketRepository.GetTicketsByProjectionID(projection.ID)
	if err != nil {
		return err
	}
	for _, ticket := range tickets {
		user, err := ts.UserRepository.GetUserByID(ticket.UserID)
		if err != nil {
			return err
		}
		user.Money += projection.Price * 1.1
		if err := ts.UserRepository.UpdateUser(user); err != nil {
			return err
		}
	}
	return nil
}
