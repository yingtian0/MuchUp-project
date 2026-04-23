package message

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/internal/usecase/dto"
	"MuchUp/app/utils"
	"context"
	"errors"
	"time"
)

type messageUsecase struct {
	messageRepo   repository.MessageRepository
	userRepo      repository.UserRepository
	messageStream repository.MessageStreamStore
}

func NewMessageUsecase(
	messageRepo repository.MessageRepository,
	userRepo repository.UserRepository,
	messageStream repository.MessageStreamStore,
) *messageUsecase {
	return &messageUsecase{
		messageRepo:   messageRepo,
		userRepo:      userRepo,
		messageStream: messageStream,
	}
}

func (u *messageUsecase) SendChatMessage(ctx context.Context, input dto.SendChatMessageInput) error {
	if input.SenderID == "" {
		return errors.New("sender id is required")
	}
	if input.RoomID == "" {
		return errors.New("room id is required")
	}
	if input.Content == "" {
		return errors.New("content is required")
	}

	user, err := u.userRepo.GetUserByID(input.SenderID)
	if err != nil {
		return err
	}
	if user.IsBlockedUsers != nil && user.IsBlockedUsers[input.SenderID] {
		return errors.New("user is blocked")
	}

	createdAt := input.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	messageID := input.MessageID
	if messageID == "" {
		messageID = utils.GenerateUUID()
	}

	message := &entity.Message{
		MessageID: messageID,
		SenderID:  input.SenderID,
		GroupID:   input.RoomID,
		Text:      utils.StringPtr(input.Content),
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	_, err = u.messageStream.AppendMessage(ctx, message)
	return err
}

func (u *messageUsecase) GetMessage(id string) (*entity.Message, error) {
	return u.messageRepo.GetMessageByID(id)
}

func (u *messageUsecase) CreateMessage(message *entity.Message) (*entity.Message, error) {
	if err := u.messageRepo.CreateMessage(message); err != nil {
		return nil, err
	}
	return message, nil
}

func (u *messageUsecase) UpdateMessage(message *entity.Message) (*entity.Message, error) {
	if err := u.messageRepo.UpdateMessage(message); err != nil {
		return nil, err
	}
	return message, nil
}

func (u *messageUsecase) DeleteMessage(id string) error {
	return u.messageRepo.DeleteMessage(id)
}
