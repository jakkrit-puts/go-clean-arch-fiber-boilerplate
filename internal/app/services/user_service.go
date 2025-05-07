package services

import (
	"go-clean-arch-fiber-boilerplate/internal/app/models"
	"go-clean-arch-fiber-boilerplate/internal/app/repositories"
)

type UserService interface {
	CreateUser(name string, email string, password string) (models.User, error)
	GetUserByID(id uint) (models.User, error)
	GetUsers() ([]models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(name string, email string, password string) (models.User, error) {
	user := models.User{Name: name, Email: email, Password: password}
	data, err := s.userRepository.Create(user)
	if err != nil {
		return models.User{}, err
	}
	return data, nil
}

func (s *userService) GetUserByID(id uint) (models.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) GetUsers() ([]models.User, error) {
	users, err := s.userRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
