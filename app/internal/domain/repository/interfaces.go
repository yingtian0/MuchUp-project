package repository

import (
	"MuchUp/app/internal/domain/entity"
	"context"
)

type UserRepository interface {
	CreateUser(user *entity.UserProfile) error
	GetUserByID(id string) (*entity.UserProfile, error)
	UpdateUser(user *entity.UserProfile) error
	DeleteUser(id string) error
	GetUsers(limit, offset int) ([]*entity.UserProfile, error)
	GetUserByEmail(email string) (*entity.UserProfile, error)
	GetUsersByRoom(roomID string) ([]*entity.UserProfile, error)
}

type RoomRepository interface {
	CreateRoom(room *entity.Room) (*entity.Room, error)
	GetRoomByID(roomID string) (*entity.Room, error)
	GetRoomsByUserID(userID string) ([]*entity.Room, error)
	UpdateRoom(room *entity.Room) (*entity.Room, error)
	DeleteRoom(roomID string) error
	AddUserToRoom(userID, roomID string) error
	FindRoomWithAvailableSlots() (*entity.Room, error)
}
type MessageRepository interface {
	CreateMessage(message *entity.Message) error
	GetMessageByID(id string) (*entity.Message, error)
	GetMessagesByUserID(userID string) ([]*entity.Message, error)
	UpdateMessage(message *entity.Message) error
	DeleteMessage(id string) error
	GetMessagesByRoom(roomID string, limit, offset int) ([]*entity.Message, error)
}

type RoomUserStore interface {
	AddConnectedUser(ctx context.Context, roomID, userID string) error
	RemoveConnectedUser(ctx context.Context, roomID, userID string) error
	ListConnectedUserIDs(ctx context.Context, roomID string) ([]string, error)
	IsConnectedUser(ctx context.Context, roomID, userID string) (bool, error)
	CountConnectedUsers(ctx context.Context, roomID string) (int64, error)
	DeleteRoomConnections(ctx context.Context, roomID string) error
}

type MessageStreamStore interface {
	AppendMessage(ctx context.Context, message *entity.Message) (string, error)
	GetRecentMessages(ctx context.Context, roomID string, count int64) ([]*entity.Message, error)
	GetMessagesAfter(ctx context.Context, roomID, lastMessageID string, count int64) ([]*entity.Message, error)
	DeleteMessageHistory(ctx context.Context, roomID string) error
}
