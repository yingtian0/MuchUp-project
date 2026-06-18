package message

import (
	"context"
	"errors"
	"time"

	usecase "MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/internal/domain/store"
	"MuchUp/app/internal/usecase/dto"

	"github.com/google/uuid"
)

type messageUsecase struct {
	messageRepo   repository.MessageRepository
	userRepo      repository.UserRepository
	messageStream store.MessageStreamStore
}

// TODO: メッセージ関連依存の初期化をこのコンストラクタに集約していく
func NewMessageUsecase(
	messageRepo repository.MessageRepository,
	userRepo repository.UserRepository,
	messageStream store.MessageStreamStore,
) usecase.MessageUsecase {
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

	_, err := u.userRepo.FindByID(ctx, input.SenderID)
	if err != nil {
		return err
	}

	createdAt := input.CreatedAt
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	messageID := input.MessageID
	if messageID == "" {
		messageID = uuid.NewString()
	}

	message := &entity.Message{
		ClientMessageID: new(messageID),
		SenderID:        entity.UserID(input.SenderID),
		RoomID:          entity.RoomID(input.RoomID),
		Text:            new(input.Content),
		CreatedAt:       createdAt,
		UpdatedAt:       createdAt,
	}

	_, err = u.messageStream.AppendMessage(ctx, message)

	return err
}

func (u *messageUsecase) GetMessage(id string) (*entity.Message, error) {
	return u.messageRepo.FindByID(context.Background(), id)
}

func (u *messageUsecase) CreateMessage(message *entity.Message) (*entity.Message, error) {
	if err := u.messageRepo.Insert(context.Background(), message); err != nil {
		return nil, err
	}

	return message, nil
}

func (u *messageUsecase) UpdateMessage(message *entity.Message) (*entity.Message, error) {
	if err := u.messageRepo.Update(context.Background(), message); err != nil {
		return nil, err
	}

	return message, nil
}

func (u *messageUsecase) DeleteMessage(id string) error {
	return u.messageRepo.Delete(context.Background(), id)
}
