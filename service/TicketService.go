package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type TicketService struct {
	TicketRepository repo.TicketRepository
}

func NewTicketService(ticketRepo repo.TicketRepository) *TicketService {
	return &TicketService{
		TicketRepository: ticketRepo,
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
