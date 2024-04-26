package model

import "time"

type Type int

// Define constants for each day of the week
const (
	Weekly Type = iota
	AllAtOnce
)

type Contract struct {
	Id              uint      `gorm:"primaryKey" json:"Id"`
	Name            string    `gorm:"not null" json:"Name"`
	SupplierId      uint      `gorm:"not null" json:"SupplierId"`
	ValidFrom       time.Time `gorm:"not null" json:"validFrom"`
	ValidUntil      time.Time `gorm:"not null" json:"validUntil"`
	DateOfSignature time.Time `gorm:"not null" json:"dateOfSignature"`
	ContractType    Type      `gorm:"not null" json:"ContractType"`
	Supplier        Supplier  `gorm:"foreignKey:SupplierId" json:"Supplier"`
}
