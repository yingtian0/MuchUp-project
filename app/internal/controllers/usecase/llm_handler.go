package usecase

import (
	"context"

	"MuchUp/app/internal/domain/entity"
)

type LLMHandler interface {
	HandleRoomCreated(ctx context.Context, room *entity.Room, owner *entity.UserProfile) error
}
