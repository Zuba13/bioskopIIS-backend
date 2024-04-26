package model

type Supplier struct {
	Id      uint   `gorm:"primaryKey" json:"Id"`
	Name    string `gorm:"not null" json:"Name"`
	Number  string `gorm:"not null" json:"Number"`
	Email   string `gorm:"not null" json:"Email"`
	Street  string `gorm:"not null" json:"Street"`
	City    string `gorm:"not null" json:"City"`
	Country string `gorm:"not null" json:"Country"`
	Phone   string `gorm:"not null" json:"Phone"`
}