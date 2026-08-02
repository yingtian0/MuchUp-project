package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type UserRepository interface {
	Insert(ctx context.Context, user *entity.UserProfile) error
	FindByID(ctx context.Context, id entity.UserID) (*entity.UserProfile, error)
	FindAll(ctx context.Context, limit, offset uint) ([]*entity.UserProfile, error)
	FindByEmail(ctx context.Context, email string) (*entity.UserProfile, error)
	FindByRoom(ctx context.Context, roomID entity.RoomID) ([]*entity.UserProfile, error)
	Update(ctx context.Context, user *entity.UserProfile) error
	Delete(ctx context.Context, id entity.UserID) error
}
