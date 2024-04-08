package model

type Hall struct {
	ID           uint `json:"id" gorm:"primaryKey"`
	NumberOfRows uint `json:"numberOfRows"`
	SeatsPerRow  uint `json:"seatsPerRow"`
}
