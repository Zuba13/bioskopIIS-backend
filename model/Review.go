package model

import "time"

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	MovieID   uint      `json:"movieId"`
	UserID    uint      `json:"userId"`
	Rating    int       `json:"rating" gorm:"check:rating >= 1 AND rating <= 10"` 
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
