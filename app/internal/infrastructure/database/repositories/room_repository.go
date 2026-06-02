package repositories

import (
	"MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/internal/infrastructure/database/mapper"
	"MuchUp/app/internal/infrastructure/database/schema"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type roomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) repository.RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) CreateRoom(room *entity.Room) (*entity.Room, error) {
	roomSchema := mapper.ToRoomSchema(room)
	if err := r.db.Create(roomSchema).Error; err != nil {
		return nil, err
	}
	if len(roomSchema.Users) > 0 {
		if err := r.db.Model(roomSchema).Association("Users").Append(roomSchema.Users); err != nil {
			return nil, err
		}
	}
	return r.GetRoomByID(string(room.ID))
}

func (r *roomRepository) GetRoomByID(roomID string) (*entity.Room, error) {
	var roomSchema schema.RoomSchema
	err := r.db.Preload("Users").First(&roomSchema, "id = ?", roomID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("room not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get room: %w", err)
	}
	return mapper.ToRoomEntity(&roomSchema), nil
}

func (r *roomRepository) GetRoomsByUserID(userID string) ([]*entity.Room, error) {
	var roomSchemas []schema.RoomSchema
	err := r.db.
		Joins("JOIN user_rooms ON user_rooms.room_id = rooms.id").
		Where("user_rooms.user_id = ?", userID).
		Preload("Users").
		Find(&roomSchemas).Error
	if err != nil {
		return nil, err
	}

	rooms := make([]*entity.Room, 0, len(roomSchemas))
	for i := range roomSchemas {
		rooms = append(rooms, mapper.ToRoomEntity(&roomSchemas[i]))
	}
	return rooms, nil
}

func (r *roomRepository) UpdateRoom(room *entity.Room) (*entity.Room, error) {
	roomSchema := mapper.ToRoomSchema(room)
	if err := r.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(roomSchema).Error; err != nil {
		return nil, err
	}
	return r.GetRoomByID(string(room.ID))
}

func (r *roomRepository) DeleteRoom(roomID string) error {
	return r.db.Delete(&schema.RoomSchema{}, "id = ?", roomID).Error
}

func (r *roomRepository) AddUserToRoom(userID, roomID string) error {
	user := schema.UserSchema{ID: userID}
	room := schema.RoomSchema{ID: roomID}
	return r.db.Model(&room).Association("Users").Append(&user)
}

func (r *roomRepository) FindRoomWithAvailableSlots() (*entity.Room, error) {
	var roomSchema schema.RoomSchema
	err := r.db.Preload("Users").
		Joins("LEFT JOIN user_rooms ON user_rooms.room_id = rooms.id").
		Where("rooms.status = ?", string(entity.RoomWaiting)).
		Group("rooms.id").
		Having("COUNT(user_rooms.user_id) < rooms.capacity").
		Order("COUNT(user_rooms.user_id) DESC").
		First(&roomSchema).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, usecase.ErrNotFound
		}
		return nil, err
	}
	return mapper.ToRoomEntity(&roomSchema), nil
}
