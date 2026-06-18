package postgres

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/sqlc"
)

type userRepository struct {
	q *sqlc.Queries
}

func NewUserRepository(db sqlc.DBTX) repository.UserRepository {
	return &userRepository{q: sqlc.New(db)}
}
func (repo *userRepository) CreateUser(_ context.Context, _ *entity.UserProfile) error {
	return ErrNotImplemented
}

func (repo *userRepository) GetUserByID(_ context.Context, _ string) (*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func (repo *userRepository) UpdateUser(_ context.Context, _ *entity.UserProfile) error {
	return ErrNotImplemented
}

func (repo *userRepository) DeleteUser(_ context.Context, _ string) error {
	return ErrNotImplemented
}

func (repo *userRepository) GetUsers(_ context.Context, _, _ int) ([]*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func (repo *userRepository) GetUserByEmail(_ context.Context, _ string) (*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func (repo *userRepository) GetUsersByRoom(_ context.Context, _ string) ([]*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}
