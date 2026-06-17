package mocks

import (
	"context"

	"MuchUp/app/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *entity.UserProfile) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
func (m *MockUserRepository) GetUserByID(ctx context.Context, id string) (*entity.UserProfile, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	user, ok := args.Get(0).(*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return user, args.Error(1)
}
func (m *MockUserRepository) UpdateUser(ctx context.Context, user *entity.UserProfile) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
func (m *MockUserRepository) DeleteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockUserRepository) GetUsers(ctx context.Context, limit, offset int) ([]*entity.UserProfile, error) {
	args := m.Called(ctx, limit, offset)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	users, ok := args.Get(0).([]*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return users, args.Error(1)
}
func (m *MockUserRepository) GetUsersByRoom(ctx context.Context, roomID string) ([]*entity.UserProfile, error) {
	args := m.Called(ctx, roomID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	users, ok := args.Get(0).([]*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return users, args.Error(1)
}
func (m *MockUserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserProfile, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	user, ok := args.Get(0).(*entity.UserProfile)
	if !ok {
		return nil, args.Error(1)
	}

	return user, args.Error(1)
}
