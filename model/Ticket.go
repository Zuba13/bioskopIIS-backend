package model

import "time"

type Ticket struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	UserID      uint      `gorm:"not null" json:"userId"`
	ProjectionID uint     `gorm:"not null" json:"projectionId"`
	RowNum      int       `gorm:"not null" json:"rowNum"`
	SeatNum     int       `gorm:"not null" json:"seatNum"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
