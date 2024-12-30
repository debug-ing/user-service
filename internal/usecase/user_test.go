package usecase

import (
	"testing"

	"github.com/debug-ing/user-service/internal/domain"
	repository "github.com/debug-ing/user-service/internal/repository/muck"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// ignore
func TestCreateSuccess(t *testing.T) {
	repo := new(repository.MockUserRepository)
	u := NewUserUsecase(repo)
	user := &domain.User{ID: 1, Username: "john", Password: "password"}
	repo.On("FindByUsername", "john").Return(nil, nil)
	repo.On("Create", mock.Anything).Return(nil)
	err := u.Create(user)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestCreateUserExist(t *testing.T) {
	repo := new(repository.MockUserRepository)
	u := NewUserUsecase(repo)
	user := &domain.User{ID: 1, Username: "john", Password: "password"}
	repo.On("FindByUsername", "john").Return(user, nil)
	repo.On("Create", mock.Anything).Return(nil)
	err := u.Create(user)
	assert.Error(t, err)
	assert.Equal(t, "username already exists", err.Error())
}

func TestFindByUsername(t *testing.T) {
	t.Skip("Skip this test")
	repo := new(repository.MockUserRepository)
	u := NewUserUsecase(repo)
	user := &domain.User{ID: 1, Username: "john", Password: "password"}
	repo.On("FindByUsername", "john").Return(user, nil)
	result, err := u.FindByUsername("john")
	assert.NoError(t, err)
	assert.Equal(t, user, result)
	repo.AssertExpectations(t)
}
