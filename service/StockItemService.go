package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type StockItemService struct {
	StockItemRepository repo.StockItemRepository
}

func NewStockItemService(menuItemRepository repo.StockItemRepository) *StockItemService {
	return &StockItemService{
		StockItemRepository: menuItemRepository,
	}
}

func (stockItemService *StockItemService) CreateStockItem(stockItem model.StockItem) (model.StockItem, error) {
	return stockItemService.StockItemRepository.Create(&stockItem)
}

func (stockItemService *StockItemService) GetAllStockItems() ([]model.StockItem, error) {
	return stockItemService.StockItemRepository.GetAll()
}

func (stockItemService *StockItemService) GetStockItemById(id uint) (model.StockItem, error) {
	return stockItemService.StockItemRepository.GetById(id)
}

func (stockItemService *StockItemService) UpdateStockItem(contract *model.StockItem) error {
	return stockItemService.StockItemRepository.Update(contract)
}

func (stockItemService *StockItemService) DeleteStockItem(id uint) error {
	return stockItemService.StockItemRepository.Delete(id)
}