package usecase

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/usecase/dto"
	"context"
)

type MessageUsecase interface {
	SendChatMessage(ctx context.Context, input dto.SendChatMessageInput) error
	GetMessage(id string) (*entity.Message, error)
	CreateMessage(message *entity.Message) (*entity.Message, error)
	UpdateMessage(message *entity.Message) (*entity.Message, error)
	DeleteMessage(id string) error
}
