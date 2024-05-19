package model

type Hall struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	Name         string `gorm:"not null" json:"name"`
	NumberOfRows uint   `gorm:"not null" json:"numberOfRows"`
	SeatsPerRow  uint   `gorm:"not null" json:"seatsPerRow"`
}
