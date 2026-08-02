package store

import (
	"context"
)

type RoomPresenceStore interface {
	AddConnectedUser(ctx context.Context, roomID, userID string) error
	RemoveConnectedUser(ctx context.Context, roomID, userID string) error
	ListConnectedUserIDs(ctx context.Context, roomID string) ([]string, error)
	IsConnectedUser(ctx context.Context, roomID, userID string) (bool, error)
	CountConnectedUsers(ctx context.Context, roomID string) (int64, error)
	DeleteRoomConnections(ctx context.Context, roomID string) error
}
