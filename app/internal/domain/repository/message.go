package repository

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type MessageRepository interface {
	Insert(ctx context.Context, message *entity.Message) error
	FindByID(ctx context.Context, id string) (*entity.Message, error)
	FindByUserID(ctx context.Context, userID string) ([]*entity.Message, error)
	FindByRoom(ctx context.Context, roomID string, limit, offset int) ([]*entity.Message, error)
	Update(ctx context.Context, message *entity.Message) error
	Delete(ctx context.Context, id string) error
}
