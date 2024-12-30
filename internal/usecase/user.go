package usecase

import (
	"errors"

	"github.com/debug-ing/user-service/internal/domain"
)

// UserUsecase is a struct to represent domain.UserUsecase
type UserUsecase struct {
	repo domain.UserRepository
}

// NewUserUsecase will create an object that represent the domain.UserUsecase
func NewUserUsecase(repo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{repo: repo}
}

// Create will create a new user
func (u *UserUsecase) Create(user *domain.User) error {
	user, err := u.repo.FindByUsername(user.Username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("username already exists")
	}
	return u.repo.Create(user)
}

// FindByUsername will find a user by username
func (u *UserUsecase) FindByUsername(username string) (*domain.User, error) {
	return u.repo.FindByUsername(username)
}
