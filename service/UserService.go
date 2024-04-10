package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (us *UserService) RegisterUser(username, email, password string) (*model.User, error) {
	hashedPassword, err := hashPassword(password)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
		Role:     "user",
	}

	err = us.UserRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUserByID(id uint) (*model.User, error) {
	return us.UserRepo.GetUserByID(id)
}

func (us *UserService) GetUserByUsername(username string) (*model.User, error) {
	return us.UserRepo.GetUserByUsername(username)
}

func (us *UserService) UpdateUser(user *model.User) error {
	return us.UserRepo.UpdateUser(user)
}

func (us *UserService) DeleteUser(id uint) error {
	user, err := us.GetUserByID(id)
	if err != nil {
		return err
	}

	return us.UserRepo.DeleteUser(user)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
