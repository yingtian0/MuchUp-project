package mapper

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/infrastructure/database/schema"
)

func ToRoomSchema(room *entity.Room) *schema.RoomSchema {
	users := make([]schema.UserSchema, 0, len(room.Members))
	for userID := range room.Members {
		users = append(users, schema.UserSchema{ID: string(userID)})
	}

	return &schema.RoomSchema{
		ID:        string(room.ID),
		Type:      string(room.Type),
		Status:    string(room.Status),
		Capacity:  room.Capacity,
		CreatedBy: string(*room.CreatedBy),
		Users:     users,
	}
}

func ToRoomEntity(roomSchema *schema.RoomSchema) *entity.Room {
	if roomSchema == nil {
		return nil
	}

	members := make(map[entity.UserID]*entity.RoomMember, len(roomSchema.Users))
	for _, userSchema := range roomSchema.Users {
		members[entity.UserID(userSchema.ID)] = &entity.RoomMember{
			UserID: entity.UserID(userSchema.ID),
			Status: entity.RoomMemberJoined,
		}
	}

	return &entity.Room{
		ID:        entity.RoomID(roomSchema.ID),
		Type:      entity.RoomType(roomSchema.Type),
		Status:    entity.RoomStatus(roomSchema.Status),
		Capacity:  roomSchema.Capacity,
		Members:   members,
		CreatedBy: new(entity.UserID(roomSchema.CreatedBy)),
		CreatedAt: roomSchema.CreatedAt,
	}
}
