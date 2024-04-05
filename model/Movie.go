package model

import "time"

type Movie struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Year        int       `gorm:"not null" json:"year"`
	Genre       string    `gorm:"not null" json:"genre"`
	Rating      float64   `gorm:"not null" json:"rating"`
	NumVotes    int       `gorm:"not null" json:"num_votes"`
	Duration    int       `gorm:"not null" json:"duration_min"`
	Image       string    `gorm:"not null" json:"image"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
