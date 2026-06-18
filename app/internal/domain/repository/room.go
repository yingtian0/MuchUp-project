package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type RoomRepository interface {
	Insert(ctx context.Context, room *entity.Room) (*entity.Room, error)
	FindByID(ctx context.Context, roomID string) (*entity.Room, error)
	FindByUserID(ctx context.Context, userID string) ([]*entity.Room, error)
	Update(ctx context.Context, room *entity.Room) (*entity.Room, error)
	Delete(ctx context.Context, roomID string) error
}
