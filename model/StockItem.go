package model

type StockItem struct {
	Id              uint    `gorm:"primaryKey;autoIncrement" json:"Id"`
	ProductId       uint    `gorm:"not null;unique" json:"ProductId"`
	Quantity        uint    `gorm:"not null" json:"Quantity"`
	Product         Product `gorm:"foreignKey:ProductId" json:"Product"`
	OptimalQuantity uint    `gorm:"not null" json:"OptimalQuantity"`
}