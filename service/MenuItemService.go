package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type MenuItemService struct {
	MenuItemRepository repo.MenuItemRepository
}

func NewMenuItemService(menuItemRepository repo.MenuItemRepository) *MenuItemService {
	return &MenuItemService{
		MenuItemRepository: menuItemRepository,
	}
}

func (menuItemService *MenuItemService) CreateMenuItem(contract model.MenuItem) (model.MenuItem, error) {
	return menuItemService.MenuItemRepository.Create(&contract)
}

func (menuItemService *MenuItemService) GetAllMenuItems() ([]model.MenuItem, error) {
	return menuItemService.MenuItemRepository.GetAll()
}

func (menuItemService *MenuItemService) GetMenuItemById(id uint) (model.MenuItem, error) {
	return menuItemService.MenuItemRepository.GetById(id)
}

func (menuItemService *MenuItemService) UpdateMenuItem(contract *model.MenuItem) error {
	return menuItemService.MenuItemRepository.Update(contract)
}

func (menuItemService *MenuItemService) DeleteMenuItem(id uint) error {
	return menuItemService.MenuItemRepository.Delete(id)
}