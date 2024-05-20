package service

import (
	"fmt"
	"time"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type StockItemService struct {
	StockItemRepository repo.StockItemRepository
	ContractService ContractService
}

func NewStockItemService(menuItemRepository repo.StockItemRepository, contractService ContractService) *StockItemService {
	return &StockItemService{
		StockItemRepository: menuItemRepository,
		ContractService: contractService,
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

func (stockItemService *StockItemService) GetStockItemByProductId(productId uint) (model.StockItem, error) {
	return stockItemService.StockItemRepository.GetById(productId)
}

func (stockItemService *StockItemService) UpdateStockItem(stockItem *model.StockItem) error {
	return stockItemService.StockItemRepository.Update(stockItem)
}

func (stockItemService *StockItemService) DeleteStockItem(id uint) error {
	return stockItemService.StockItemRepository.Delete(id)
}

func (stockItemService *StockItemService) DailyTaskForDelivery() {
	
	// Your task logic here
	atOnceContracts, err := stockItemService.ContractService.GetTodayAtOnceContracts()
	if err != nil {
		fmt.Println("Neuspjesno dobavljanje ugovora!")
		return
	}

	weeklyContracts, err := stockItemService.ContractService.GetTodayyWeeklyContracts()
	if err != nil {
		fmt.Println("Neuspjesno dobavljanje sedmicnih ugovora!")
		return
	}

	contracts := append(weeklyContracts,atOnceContracts...)

	for _, contract := range contracts {
		for _, item := range contract.ContractItems {
			stock, err := stockItemService.GetStockItemByProductId(item.ProductId)
			if err != nil {
				fmt.Printf("Failed to get stock item %d from stocks %v\n", item.Id, err)
			}
			stock.Quantity += item.Quantity // Example update
			if err := stockItemService.UpdateStockItem(&stock); err != nil {
				fmt.Printf("Failed to update stock item %d for contract %d: %v\n", item.Id, contract.Id, err)
			}
		}
	}
	fmt.Println("Contract products for today already delivered at: ", time.Now())
}

