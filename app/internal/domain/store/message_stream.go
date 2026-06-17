package store

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type MessageStreamStore interface {
	AppendMessage(ctx context.Context, message *entity.Message) (string, error)
	GetRecentMessages(ctx context.Context, roomID string, count int64) ([]*entity.Message, error)
	GetMessagesAfter(ctx context.Context, roomID, lastMessageID string, count int64) ([]*entity.Message, error)
	DeleteMessageHistory(ctx context.Context, roomID string) error
}
