package postgres

import (
	"MuchUp/app/internal/domain/repository"
)

type userRepository struct {
	repository.UserRepository
}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}
