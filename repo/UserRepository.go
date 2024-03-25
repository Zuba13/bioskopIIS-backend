package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func NewuserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DatabaseConnection: db};
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	return ur.DatabaseConnection.Create(user).Error
}


func (ur *UserRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := ur.DatabaseConnection.First(&user, id).Error; err != nil {
			return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := ur.DatabaseConnection.Where("username = ?", username).First(&user).Error; err != nil {
			return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) UpdateUser(user *model.User) error {
	return ur.DatabaseConnection.Save(user).Error
}

func (ur *UserRepository) DeleteUser(user *model.User) error {
	return ur.DatabaseConnection.Delete(user).Error
}