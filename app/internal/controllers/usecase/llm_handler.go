package usecase

import (
	"MuchUp/app/internal/domain/entity"
	"context"
)

type LLMHandler interface {
	HandleRoomCreated(ctx context.Context, room *entity.Room, owner *entity.UserProfile) error
}
