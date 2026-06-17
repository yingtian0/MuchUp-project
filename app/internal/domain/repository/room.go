package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error)
	GetRoomByID(ctx context.Context, roomID string) (*entity.Room, error)
	GetRoomsByUserID(ctx context.Context, userID string) ([]*entity.Room, error)
	UpdateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error)
	DeleteRoom(ctx context.Context, roomID string) error
	AddUserToRoom(ctx context.Context, userID, roomID string) error
	FindRoomWithAvailableSlots(ctx context.Context) (*entity.Room, error)
}
