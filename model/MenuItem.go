package model

type MenuItem struct {
	Id        uint    `gorm:"primaryKey;autoIncrement" json:"Id"`
	MenuId    uint    `gorm:"not null;index:idx_menu_product,unique" json:"MenuId"`
	ProductId uint    `gorm:"not null;index:idx_menu_product,unique" json:"ProductId"`
	Price     float32 `gorm:"not null" json:"Price"`
	Product   Product `gorm:"foreignKey:ProductId" json:"Product"`
}