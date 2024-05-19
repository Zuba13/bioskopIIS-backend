package model

import "time"

type Menu struct {
	Id              uint           `gorm:"primaryKey;autoIncrement" json:"Id"`
	Name            string         `gorm:"not null" json:"Name"`
	ValidFrom       time.Time      `gorm:"not null" json:"validFrom"`
	ValidUntil      time.Time      `gorm:"not null" json:"validUntil"`
	MenuItems   	[]MenuItem 	   `gorm:"foreignKey:MenuId" json:"MenuItems"`
}