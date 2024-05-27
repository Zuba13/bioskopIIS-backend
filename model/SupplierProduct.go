package model

type SupplierProduct struct {
	SupplierId uint     `gorm:"primaryKey" json:"SupplierId"`
	ProductId  uint     `gorm:"primaryKey" json:"ProductId"`
	Price      float32  `gorm:"not null" json:"Price"`
	Product    Product  `gorm:"foreignKey:ProductId" json:"Product"`
	Supplier   Supplier `gorm:"foreignKey:SupplierId" json:"Supplier"`
}
