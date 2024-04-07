package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type TicketRepository struct {
	DB *gorm.DB
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		DB: db,
	}
}

func (tr *TicketRepository) CreateTicket(ticket model.Ticket) (*model.Ticket, error) {
	if err := tr.DB.Create(&ticket).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (tr *TicketRepository) GetAllTickets() ([]*model.Ticket, error) {
	var tickets []*model.Ticket
	if err := tr.DB.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (tr *TicketRepository) GetTicketByID(id uint) (*model.Ticket, error) {
	var ticket model.Ticket
	if err := tr.DB.First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (tr *TicketRepository) UpdateTicket(ticket *model.Ticket) error {
	if err := tr.DB.Save(ticket).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TicketRepository) DeleteTicket(id uint) error {
	if err := tr.DB.Delete(&model.Ticket{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (tr *TicketRepository) GetTicketsByUserID(userID uint) ([]*model.Ticket, error) {
	var tickets []*model.Ticket
	if err := tr.DB.Where("user_id = ?", userID).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}
