package service

import (
	model "bioskop.com/projekat/bioskopIIS-backend/model"
	repo "bioskop.com/projekat/bioskopIIS-backend/repo"
)

type MenuService struct {
	MenuRepository repo.MenuRepository
}

func NewMenuService(contractRepository repo.MenuRepository) *MenuService {
	return &MenuService{
		MenuRepository: contractRepository,
	}
}

func (menuService *MenuService) CreateMenu(contract model.Menu) (model.Menu, error) {
	return menuService.MenuRepository.Create(&contract)
}

func (menuService *MenuService) GetAllMenus() ([]model.Menu, error) {
	return menuService.MenuRepository.GetAll()
}

func (menuService *MenuService) GetMenuById(id uint) (model.Menu, error) {
	return menuService.MenuRepository.GetById(id)
}

func (menuService *MenuService) UpdateMenu(contract *model.Menu) error {
	return menuService.MenuRepository.Update(contract)
}

func (menuService *MenuService) DeleteMenu(id uint) error {
	return menuService.MenuRepository.Delete(id)
}