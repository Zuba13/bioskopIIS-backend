package model

import "time"

type OrderSuggestion struct {
	Id         uint      `gorm:"primaryKey;autoIncrement" json:"Id"`
	SupplierId uint      `gorm:"not null" json:"SupplierId"`
	ProductId  uint      `gorm:"not null" json:"ProductId"`
	Quantity   uint      `gorm:"not null" json:"Quantity"`
	CreatedAt  time.Time `json:"created_at"`
}