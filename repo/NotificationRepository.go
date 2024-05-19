package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	DatabaseConnection *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{
		DatabaseConnection: db,
	}
}

func (nr *NotificationRepository) GetNotificationByID(id uint) (*model.ProjectionCanceledNotification, error) {
	var notification model.ProjectionCanceledNotification
	if err := nr.DatabaseConnection.First(&notification, id).Error; err != nil {
		return nil, err
	}
	return &notification, nil
}

func (nr *NotificationRepository) CreateNotification(notification *model.ProjectionCanceledNotification) (*model.ProjectionCanceledNotification, error) {
	if err := nr.DatabaseConnection.Create(notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (nr *NotificationRepository) UpdateNotification(notification *model.ProjectionCanceledNotification) (*model.ProjectionCanceledNotification, error) {
	if err := nr.DatabaseConnection.Save(notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (nr *NotificationRepository) GetCanceledProjectionNotifications(userId uint) ([]model.ProjectionCanceledNotification, error) {
	var notifications []model.ProjectionCanceledNotification
	if err := nr.DatabaseConnection.Where("user_id = ? AND is_read = ?", userId, false).Find(&notifications).Error; err != nil {
		return nil, err
	}
	return notifications, nil
}
