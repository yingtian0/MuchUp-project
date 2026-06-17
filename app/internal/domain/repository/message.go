package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *entity.Message) error
	GetMessageByID(ctx context.Context, id string) (*entity.Message, error)
	GetMessagesByUserID(ctx context.Context, userID string) ([]*entity.Message, error)
	UpdateMessage(ctx context.Context, message *entity.Message) error
	DeleteMessage(ctx context.Context, id string) error
	GetMessagesByRoom(ctx context.Context, roomID string, limit, offset int) ([]*entity.Message, error)
}
