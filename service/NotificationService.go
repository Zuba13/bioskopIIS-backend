package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

type NotificationService struct {
	NotificationRepo *repo.NotificationRepository
	TicketRepo       *repo.TicketRepository
}

func NewNotificationService(nr *repo.NotificationRepository, tr *repo.TicketRepository) *NotificationService {
	return &NotificationService{
		NotificationRepo: nr,
		TicketRepo:       tr,
	}
}

func (ns *NotificationService) GetCanceledProjectionNotifications(userID uint) ([]model.ProjectionCanceledNotification, error) {
	notifications, err := ns.NotificationRepo.GetCanceledProjectionNotifications(userID)
	if err != nil {
		return nil, err
	}

	for i := range notifications {
		notifications[i].MarkAsRead()

		if _, err := ns.NotificationRepo.UpdateNotification(&notifications[i]); err != nil {
			return nil, err
		}
	}

	return notifications, nil
}

func (ns *NotificationService) NotifyUsersAboutCancellation(projectionID uint) error {
	tickets, err := ns.TicketRepo.GetTicketsByProjectionID(projectionID)
	if err != nil {
		return err
	}

	for _, ticket := range tickets {
		notification := &model.ProjectionCanceledNotification{
			Notification: model.Notification{
				UserID: ticket.UserID,
				Type:   model.ProjectionCanceled,
				IsRead: false,
				Text:   "We are sorry to inform you that the projection you have a ticket for has been canceled. We will refund your money as soon as possible.",
			},
			TicketID: ticket.ID,
		}

		if _, err := ns.NotificationRepo.CreateNotification(notification); err != nil {
			return err
		}
	}

	return nil
}
