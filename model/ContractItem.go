package model

type ContractItem struct {
	Id         uint    `gorm:"primaryKey;autoIncrement" json:"Id"`
	ContractId uint    `gorm:"foreignKey" json:"ContractId"`
	ProductId  uint    `gorm:"not null" json:"ProductId"`
	Quantity   uint    `gorm:"not null" json:"Quantity"`
	Price      float32 `gorm:"not null" json:"Price"`
	Product    Product `gorm:"foreignKey:ProductId" json:"Product"`
}

type ContractItemDTO struct {
	ProductId uint    `json:"productId"`
	Quantity  uint    `json:"quantity"`
	Price     float32 `json:"price"`
}