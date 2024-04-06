package model

import "time"

type Projection struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	MovieID    uint      `json:"movieId"`
	TheatreID  uint      `json:"theatreId"`
	Date       time.Time `json:"date"`
	Time       string    `json:"time"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
