package model

type Supplier struct {
	Id      uint   `gorm:"primaryKey;autoIncrement" json:"Id"`
	Name    string `gorm:"not null" json:"Name"`
	Email   string `gorm:"not null" json:"Email"`
	Street  string `gorm:"not null" json:"Street"`
	City    string `gorm:"not null" json:"City"`
	Country string `gorm:"not null" json:"Country"`
	Phone   string `gorm:"not null" json:"Phone"`
}