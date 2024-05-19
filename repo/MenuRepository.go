package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type MenuRepository struct {
	DatabaseConnection *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *MenuRepository {
	return &MenuRepository{DatabaseConnection: db}
}

func (cr *MenuRepository) Create(menu *model.Menu) (model.Menu, error) {
	if err := cr.DatabaseConnection.Create(menu).Error; err != nil {
		return model.Menu{}, err
	}
	return *menu, nil
}

func (cr *MenuRepository) GetAll() ([]model.Menu, error) {
	var menus []model.Menu
	if err := cr.DatabaseConnection.Preload("MenuItems.Product").Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (cr *MenuRepository) GetById(id uint) (model.Menu, error) {
	var menu model.Menu
	if err := cr.DatabaseConnection.Preload("MenuItems.Product").First(&menu, id).Error; err != nil {
		return model.Menu{}, err
	}
	return menu, nil
}

func (cr *MenuRepository) Update(menu *model.Menu) error {
	return cr.DatabaseConnection.Save(menu).Error
}

func (cr *MenuRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.Menu{}, id).Error
}