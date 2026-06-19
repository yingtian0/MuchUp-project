package mocks

import (
	"context"

	"MuchUp/app/internal/domain/entity"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Insert(ctx context.Context, user *entity.UserProfile) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
func (m *MockUserRepository) FindByID(ctx context.Context, id entity.UserID) (*entity.UserProfile, error) {
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
func (m *MockUserRepository) Update(ctx context.Context, user *entity.UserProfile) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}
func (m *MockUserRepository) Delete(ctx context.Context, id entity.UserID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
func (m *MockUserRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.UserProfile, error) {
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
func (m *MockUserRepository) FindByRoom(ctx context.Context, roomID entity.RoomID) ([]*entity.UserProfile, error) {
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
func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*entity.UserProfile, error) {
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
