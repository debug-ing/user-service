package repository

import (
	domain "github.com/debug-ing/user-service/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository will create an object that represent the domain.UserRepository
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// Create will create a new user
func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// FindByUsername will find a user by username
func (r *userRepository) FindByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
