package usecase

import (
	"MuchUp/app/internal/domain/entity"
)

type GroupUsecase interface {
	FindOrCreateGroupForUser(user *entity.User) (*entity.ChatGroup, error)
	AddUserToGroup(userID, groupID string) error
}
