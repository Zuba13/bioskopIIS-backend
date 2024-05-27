package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type OrderSuggestionRepository struct {
	DatabaseConnection *gorm.DB
}

func NewOrderSuggestionRepository(db *gorm.DB) *OrderSuggestionRepository {
	return &OrderSuggestionRepository{DatabaseConnection: db}
}

func (cr *OrderSuggestionRepository) Create(orderSuggestion *model.OrderSuggestion) (model.OrderSuggestion, error) {
	if err := cr.DatabaseConnection.Create(orderSuggestion).Error; err != nil {
		return model.OrderSuggestion{}, err
	}
	return *orderSuggestion, nil
}

func (cr *OrderSuggestionRepository) GetAll() ([]model.OrderSuggestion, error) {
	var orderSuggestions []model.OrderSuggestion
	if err := cr.DatabaseConnection.Find(&orderSuggestions).Error; err != nil {
		return nil, err
	}
	return orderSuggestions, nil
}

func (cr *OrderSuggestionRepository) GetById(id uint) (model.OrderSuggestion, error) {
	var orderSuggestion model.OrderSuggestion
	if err := cr.DatabaseConnection.First(&orderSuggestion, id).Error; err != nil {
		return model.OrderSuggestion{}, err
	}
	return orderSuggestion, nil
}

func (cr *OrderSuggestionRepository) Update(orderSuggestion *model.OrderSuggestion) error {
	return cr.DatabaseConnection.Save(orderSuggestion).Error
}

func (cr *OrderSuggestionRepository) Delete(id uint) error {
	return cr.DatabaseConnection.Delete(&model.OrderSuggestion{}, id).Error
}