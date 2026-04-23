package usecase

import (
	"MuchUp/app/internal/domain/entity"
	"context"
)

type LLMHandler interface {
	HandleRoomCreated(ctx context.Context, group *entity.ChatGroup, owner *entity.User) error
}
