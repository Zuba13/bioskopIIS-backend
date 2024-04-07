package model

import "time"

type Projection struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	MovieID    uint      `json:"movieId"`
	HallId     uint      `json:"hallId"`
	Date       time.Time `json:"date"`
	Time       string    `json:"time"`
	Price      float64   `json:"price"` 
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
