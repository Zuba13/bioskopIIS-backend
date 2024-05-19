package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type MenuItemRepository struct {
	DatabaseConnection *gorm.DB
}

func NewMenuItemRepository(db *gorm.DB) *MenuItemRepository {
	return &MenuItemRepository{DatabaseConnection: db}
}

func (cr *MenuItemRepository) Create(menuItem *model.MenuItem) (model.MenuItem, error) {
	if err := cr.DatabaseConnection.Create(menuItem).Error; err != nil {
		return model.MenuItem{}, err
	}
	return *menuItem, nil
}

func (cr *MenuItemRepository) GetAll() ([]model.MenuItem, error) {
	var menuItems []model.MenuItem
	if err := cr.DatabaseConnection.Preload("Product").Find(&menuItems).Error; err != nil {
		return nil, err
	}
	return menuItems, nil
}

func (cr *MenuItemRepository) GetById(id uint) (model.MenuItem, error) {
	var menuItem model.MenuItem
	if err := cr.DatabaseConnection.Preload("Product").First(&menuItem, id).Error; err != nil {
		return model.MenuItem{}, err
	}
	return menuItem, nil
}

func (cr *MenuItemRepository) Update(menuItem *model.MenuItem) error {
	return cr.DatabaseConnection.Save(menuItem).Error
}

func (cr *MenuItemRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.MenuItem{}, id).Error
}