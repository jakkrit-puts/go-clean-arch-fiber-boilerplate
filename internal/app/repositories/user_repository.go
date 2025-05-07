package repositories

import (
	"go-clean-arch-fiber-boilerplate/internal/app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByID(id uint) (models.User, error)
	FindAll() ([]models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user models.User) (models.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) FindByID(id uint) (models.User, error) {
	var user models.User

	tx := r.db.First(&user, id)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}

	return user, nil
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
