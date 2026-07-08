package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type MessageRepository interface {
	Insert(ctx context.Context, message *entity.Message) error
	FindByID(ctx context.Context, id entity.MessageID) (*entity.Message, error)
	FindByUserID(ctx context.Context, userID entity.UserID, limit, offset int) ([]*entity.Message, error)
	FindByRoomID(ctx context.Context, roomID entity.RoomID, limit, offset int) ([]*entity.Message, error)
	Update(ctx context.Context, message *entity.Message) error
	Delete(ctx context.Context, id entity.MessageID) error
}
