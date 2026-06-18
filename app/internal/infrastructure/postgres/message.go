package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/sqlc"
)

type messageRepository struct {
	q *sqlc.Queries
}

func NewMessageRepository(db sqlc.DBTX) repository.MessageRepository {
	return &messageRepository{q: sqlc.New(db)}
}

func (repo *messageRepository) Insert(_ context.Context, _ *entity.Message) error {
	return ErrNotImplemented
}

func (repo *messageRepository) FindByID(_ context.Context, _ string) (*entity.Message, error) {
	return nil, ErrNotImplemented
}

func (repo *messageRepository) FindByUserID(_ context.Context, _ string) ([]*entity.Message, error) {
	return nil, ErrNotImplemented
}

func (repo *messageRepository) Update(_ context.Context, _ *entity.Message) error {
	return ErrNotImplemented
}

func (repo *messageRepository) Delete(_ context.Context, _ string) error {
	return ErrNotImplemented
}

func (repo *messageRepository) FindByRoom(_ context.Context, _ string, _, _ int) ([]*entity.Message, error) {
	return nil, ErrNotImplemented
}
