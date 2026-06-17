package postgres

import (
	"MuchUp/app/internal/domain/repository"
)

type messageRepository struct {
	repository.MessageRepository
}

func NewMessageRepository() repository.MessageRepository {
	return &messageRepository{}
}
