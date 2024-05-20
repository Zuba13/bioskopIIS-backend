package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type StockItemRepository struct {
	DatabaseConnection *gorm.DB
}

func NewStockItemRepository(db *gorm.DB) *StockItemRepository {
	return &StockItemRepository{DatabaseConnection: db}
}

func (cr *StockItemRepository) Create(stockItem *model.StockItem) (model.StockItem, error) {
	if err := cr.DatabaseConnection.Create(stockItem).Error; err != nil {
		return model.StockItem{}, err
	}
	return *stockItem, nil
}

func (cr *StockItemRepository) GetAll() ([]model.StockItem, error) {
	var stockItems []model.StockItem
	if err := cr.DatabaseConnection.Preload("Product").Find(&stockItems).Error; err != nil {
		return nil, err
	}
	return stockItems, nil
}

func (cr *StockItemRepository) GetById(id uint) (model.StockItem, error) {
	var stockItem model.StockItem
	if err := cr.DatabaseConnection.Preload("Product").First(&stockItem, id).Error; err != nil {
		return model.StockItem{}, err
	}
	return stockItem, nil
}

func (cr *StockItemRepository) Update(stockItem *model.StockItem) error {
	return cr.DatabaseConnection.Save(stockItem).Error
}

func (cr *StockItemRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.StockItem{}, id).Error
}

func (cr *StockItemRepository) GetByProductId(productId uint) (model.StockItem, error) {
	var stockItem model.StockItem
	if err := cr.DatabaseConnection.Preload("Product").First(&stockItem, productId).Error; err != nil {
		return model.StockItem{}, err
	}
	return stockItem, nil
}