package muck

import (
	"github.com/debug-ing/user-service/internal/domain"

	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock type for domain.UserRepository
type MockUserRepository struct {
	mock.Mock
}

// Create is a mock function for domain.UserRepository.Create
func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// FindByUsername is a mock function for domain.UserRepository.FindByUsername
func (m *MockUserRepository) FindByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}
