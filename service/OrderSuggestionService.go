package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type OrderSuggestionService struct {
	OrderSuggestionRepository repo.OrderSuggestionRepository
}

func NewOrderSuggestionService(menuItemRepository repo.OrderSuggestionRepository) *OrderSuggestionService {
	return &OrderSuggestionService{
		OrderSuggestionRepository: menuItemRepository,
	}
}

func (menuItemService *OrderSuggestionService) CreateOrderSuggestion(orderSuggestion model.OrderSuggestion) (model.OrderSuggestion, error) {
	return menuItemService.OrderSuggestionRepository.Create(&orderSuggestion)
}

func (menuItemService *OrderSuggestionService) GetAllOrderSuggestions() ([]model.OrderSuggestion, error) {
	return menuItemService.OrderSuggestionRepository.GetAll()
}

func (menuItemService *OrderSuggestionService) GetOrderSuggestionById(id uint) (model.OrderSuggestion, error) {
	return menuItemService.OrderSuggestionRepository.GetById(id)
}

func (menuItemService *OrderSuggestionService) UpdateOrderSuggestion(orderSuggestion *model.OrderSuggestion) error {
	return menuItemService.OrderSuggestionRepository.Update(orderSuggestion)
}

func (menuItemService *OrderSuggestionService) DeleteOrderSuggestion(id uint) error {
	return menuItemService.OrderSuggestionRepository.Delete(id)
}