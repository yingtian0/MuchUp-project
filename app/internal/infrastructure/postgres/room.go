package postgres

import (
	"MuchUp/app/internal/domain/repository"
)

type roomRepository struct {
	repository.RoomRepository
}

func NewRoomRepository() repository.RoomRepository {
	return &roomRepository{}
}
