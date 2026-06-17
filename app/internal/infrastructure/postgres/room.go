package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type roomRepository struct {
	db *pgxpool.Pool
}

func NewRoomRepository(db *pgxpool.Pool) repository.RoomRepository {
	return &roomRepository{db: db}
}

func (repo *roomRepository) CreateRoom(_ context.Context, _ *entity.Room) (*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) GetRoomByID(_ context.Context, _ string) (*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) GetRoomsByUserID(_ context.Context, _ string) ([]*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) UpdateRoom(_ context.Context, _ *entity.Room) (*entity.Room, error) {
	return nil, ErrNotImplemented
}

func (repo *roomRepository) DeleteRoom(_ context.Context, _ string) error {
	return ErrNotImplemented
}

func (repo *roomRepository) AddUserToRoom(_ context.Context, _, _ string) error {
	return ErrNotImplemented
}

func (repo *roomRepository) FindRoomWithAvailableSlots(_ context.Context) (*entity.Room, error) {
	return nil, ErrNotImplemented
}
