package model

type TheatreInfo struct {
	ID          uint `gorm:"primaryKey"`
	OpeningHour int  `gorm:"not null"`
	ClosingHour int  `gorm:"not null"`
	// Add other fields as needed
}
