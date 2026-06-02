package mocks

import (
	"MuchUp/app/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(user *entity.UserProfile) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) GetUserByID(id string) (*entity.UserProfile, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	user, ok := args.Get(0).(*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return user, args.Error(1)
}
func (m *MockUserRepository) UpdateUser(user *entity.UserProfile) error {
	args := m.Called(user)
	return args.Error(0)
}
func (m *MockUserRepository) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}
func (m *MockUserRepository) GetUsers(limit, offset int) ([]*entity.UserProfile, error) {
	args := m.Called(limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	users, ok := args.Get(0).([]*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return users, args.Error(1)
}
func (m *MockUserRepository) GetUsersByRoom(roomID string) ([]*entity.UserProfile, error) {
	args := m.Called(roomID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	users, ok := args.Get(0).([]*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return users, args.Error(1)
}
func (m *MockUserRepository) GetUserByEmail(email string) (*entity.UserProfile, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	user, ok := args.Get(0).(*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return user, args.Error(1)
}
