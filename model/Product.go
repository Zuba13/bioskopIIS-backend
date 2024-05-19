package model

type Product struct {
	Id       uint   `gorm:"primaryKey;autoIncrement" json:"Id"`
	Name     string `gorm:"not null" json:"Name"`
	Producer string `gorm:"not null" json:"Producer"`
}