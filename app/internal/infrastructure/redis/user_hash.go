package redis

import (
	"MuchUp/app/internal/domain/repository"
	"context"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type RoomUserHashStore struct {
	client goredis.Cmdable
}

var _ repository.RoomUserStore = (*RoomUserHashStore)(nil)

func NewRoomUserHashStore(client goredis.Cmdable) *RoomUserHashStore {
	return &RoomUserHashStore{client: client}
}

func (s *RoomUserHashStore) AddConnectedUser(ctx context.Context, roomID, userID string) error {
	return s.client.HSet(ctx, roomUsersKey(roomID), userID, strconv.FormatInt(time.Now().UnixMilli(), 10)).Err()
}

func (s *RoomUserHashStore) RemoveConnectedUser(ctx context.Context, roomID, userID string) error {
	return s.client.HDel(ctx, roomUsersKey(roomID), userID).Err()
}

func (s *RoomUserHashStore) ListConnectedUserIDs(ctx context.Context, roomID string) ([]string, error) {
	return s.client.HKeys(ctx, roomUsersKey(roomID)).Result()
}

func (s *RoomUserHashStore) IsConnectedUser(ctx context.Context, roomID, userID string) (bool, error) {
	return s.client.HExists(ctx, roomUsersKey(roomID), userID).Result()
}

func (s *RoomUserHashStore) CountConnectedUsers(ctx context.Context, roomID string) (int64, error) {
	return s.client.HLen(ctx, roomUsersKey(roomID)).Result()
}

func (s *RoomUserHashStore) DeleteRoomConnections(ctx context.Context, roomID string) error {
	return s.client.Del(ctx, roomUsersKey(roomID)).Err()
}

func roomUsersKey(roomID string) string {
	return "room:" + roomID + ":users"
}
