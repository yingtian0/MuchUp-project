package group

import (
	usecase "MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/pkg/logger"
	"context"
)

type groupUsecase struct {
	groupRepo repository.ChatGroupRepository
	llm       usecase.LLMHandler
	logger    logger.Logger
}

func NewGroupUsecase(
	groupRepo repository.ChatGroupRepository,
	llm usecase.LLMHandler,
	logger logger.Logger,
) usecase.GroupUsecase {
	return &groupUsecase{
		groupRepo: groupRepo,
		llm:       llm,
		logger:    logger,
	}
}
func (u *groupUsecase) FindOrCreateGroupForUser(user *entity.User) (*entity.ChatGroup, error) {
	group, err := u.groupRepo.FindGroupWithAvailableSlots()
	if err != nil {
		u.logger.WithField("user_id", user.ID).Info("No available group found, creating a new one.")
		newGroup := entity.NewChatGroup("Automatic Group", 6, *user)
		createdGroup, err := u.groupRepo.CreateGroup(newGroup)
		if err != nil {
			u.logger.WithError(err).Error("Failed to create a new group")
			return nil, err
		}
		if err := u.PublishRoomCreated(context.Background(), createdGroup, user); err != nil {
			u.logger.WithError(err).WithField("group_id", createdGroup.ID).Warn("Failed to publish room created event to llm service")
		}
		return createdGroup, nil
	}
	if err := u.groupRepo.AddUserToGroup(user.ID, group.ID); err != nil {
		u.logger.WithError(err).WithField("group_id", group.ID).Error("Failed to add user to group")
		return nil, err
	}
	group.Members = append(group.Members, *user)
	u.logger.WithFields(map[string]interface{}{
		"user_id":  user.ID,
		"group_id": group.ID,
	}).Info("User added to existing group")
	return group, nil
}
func (u *groupUsecase) AddUserToGroup(userID, groupID string) error {
	return u.groupRepo.AddUserToGroup(userID, groupID)
}

func (u *groupUsecase) PublishRoomCreated(ctx context.Context, group *entity.ChatGroup, owner *entity.User) error {
	if u.llm == nil {
		return nil
	}
	return u.llm.HandleRoomCreated(ctx, group, owner)
}
