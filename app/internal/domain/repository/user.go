package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.UserProfile) error
	GetUserByID(ctx context.Context, id string) (*entity.UserProfile, error)
	UpdateUser(ctx context.Context, user *entity.UserProfile) error
	DeleteUser(ctx context.Context, id string) error
	GetUsers(ctx context.Context, limit, offset int) ([]*entity.UserProfile, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.UserProfile, error)
	GetUsersByRoom(ctx context.Context, roomID string) ([]*entity.UserProfile, error)
}
