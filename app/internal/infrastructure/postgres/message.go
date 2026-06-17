package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type messageRepository struct {
	db *pgxpool.Pool
}

func NewMessageRepository(db *pgxpool.Pool) repository.MessageRepository {
	return &messageRepository{db: db}
}

func (repo *messageRepository) CreateMessage(_ context.Context, _ *entity.Message) error {
	return ErrNotImplemented
}

func (repo *messageRepository) GetMessageByID(_ context.Context, _ string) (*entity.Message, error) {
	return nil, ErrNotImplemented
}

func (repo *messageRepository) GetMessagesByUserID(_ context.Context, _ string) ([]*entity.Message, error) {
	return nil, ErrNotImplemented
}

func (repo *messageRepository) UpdateMessage(_ context.Context, _ *entity.Message) error {
	return ErrNotImplemented
}

func (repo *messageRepository) DeleteMessage(_ context.Context, _ string) error {
	return ErrNotImplemented
}

func (repo *messageRepository) GetMessagesByRoom(_ context.Context, _ string, _, _ int) ([]*entity.Message, error) {
	return nil, ErrNotImplemented
}
