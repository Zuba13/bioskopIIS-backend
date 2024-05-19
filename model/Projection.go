package model

import "time"

type Projection struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	MovieID    uint      `gorm:"not null" json:"movieId"`
	HallId     uint      `gorm:"not null" json:"hallId"`
	Timeslot   Timeslot  `gorm:"embedded;embeddedPrefix:timeslot_" json:"timeslot"`
	Price      float64   `gorm:"not null" json:"price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsCanceled bool      `gorm:"not null" json:"is_canceled"`
}
