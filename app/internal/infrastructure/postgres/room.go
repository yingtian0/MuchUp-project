package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/sqlc"
)

type roomRepository struct {
	q *sqlc.Queries
}

func NewRoomRepository(db sqlc.DBTX) repository.RoomRepository {
	return &roomRepository{q: sqlc.New(db)}
}

func (repo *roomRepository) Insert(_ context.Context, _ *entity.Room) (*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) FindByID(_ context.Context, _ string) (*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) FindByUserID(_ context.Context, _ string) ([]*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) Update(_ context.Context, _ *entity.Room) (*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) Delete(_ context.Context, _ string) error {
	return ErrNotImplemented
}
