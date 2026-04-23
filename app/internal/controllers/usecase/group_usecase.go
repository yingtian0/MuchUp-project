package usecase

import (
	"MuchUp/app/internal/domain/entity"
	"context"
)

type GroupUsecase interface {
	FindOrCreateGroupForUser(user *entity.User) (*entity.ChatGroup, error)
	AddUserToGroup(userID, groupID string) error
	PublishRoomCreated(ctx context.Context, group *entity.ChatGroup, owner *entity.User) error
}
