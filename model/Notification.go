package model

type NotificationType string

const (
	ProjectionCanceled NotificationType = "ProjectionCanceled"
)

type Notification struct {
	ID     uint             `gorm:"primary_key" json:"id"`
	UserID uint             `gorm:"not null" json:"userId"`
	Type   NotificationType `gorm:"not null" json:"type"`
	IsRead bool             `gorm:"not null" json:"isRead"`
	Text   string           ` json:"text"`
}

func (n *Notification) MarkAsRead() {
	n.IsRead = true
}

type ProjectionCanceledNotification struct {
	Notification
	TicketID uint `gorm:"not null" json:"ticketId"`
}
