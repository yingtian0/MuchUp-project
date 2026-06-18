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
func (repo *userRepository) Insert(_ context.Context, _ *entity.UserProfile) error {
	return ErrNotImplemented
}

func (repo *userRepository) FindByID(_ context.Context, _ string) (*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func (repo *userRepository) Update(_ context.Context, _ *entity.UserProfile) error {
	return ErrNotImplemented
}

func (repo *userRepository) Delete(_ context.Context, _ string) error {
	return ErrNotImplemented
}

func (repo *userRepository) FindAll(_ context.Context, _, _ int) ([]*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func (repo *userRepository) FindByEmail(_ context.Context, _ string) (*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}

func (repo *userRepository) FindByRoom(_ context.Context, _ string) ([]*entity.UserProfile, error) {
	return nil, ErrNotImplemented
}
