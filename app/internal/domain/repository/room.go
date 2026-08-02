package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type RoomRepository interface {
	Insert(ctx context.Context, room *entity.Room) error
	FindByID(ctx context.Context, roomID entity.RoomID) (*entity.Room, error)
	FindByUserID(ctx context.Context, userID entity.UserID) ([]*entity.Room, error)
	Update(ctx context.Context, room *entity.Room) error
	Delete(ctx context.Context, roomID entity.RoomID) error
}
