package model

type ContractItem struct {
	Id         uint    `gorm:"primaryKey;autoIncrement" json:"Id"`
	ContractId uint    `gorm:"foreignKey" json:"ContractId"`
	Name       string  `gorm:"not null" json:"Name"`
	Quantity   uint    `gorm:"not null" json:"Quantity"`
	Price      float32 `gorm:"not null" json:"Price"`
}

type ContractItemDTO struct {
	Name     string  `json:"name"`
	Quantity uint    `json:"quantity"`
	Price    float32 `json:"price"`
}