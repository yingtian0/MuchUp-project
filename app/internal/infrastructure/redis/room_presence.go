package redis

import (
	"context"
	"strconv"
	"time"

	"MuchUp/app/internal/domain/store"

	goredis "github.com/redis/go-redis/v9"
)

type roomPresenceStore struct {
	client goredis.Cmdable
}

var _ store.RoomPresenceStore = (*roomPresenceStore)(nil)

func NewRoomPresenceStore(client goredis.Cmdable) store.RoomPresenceStore {
	return &roomPresenceStore{client: client}
}

func (s *roomPresenceStore) AddConnectedUser(ctx context.Context, roomID, userID string) error {
	return s.client.HSet(ctx, roomUsersKey(roomID), userID, strconv.FormatInt(time.Now().UnixMilli(), 10)).Err()
}

func (s *roomPresenceStore) RemoveConnectedUser(ctx context.Context, roomID, userID string) error {
	return s.client.HDel(ctx, roomUsersKey(roomID), userID).Err()
}

func (s *roomPresenceStore) ListConnectedUserIDs(ctx context.Context, roomID string) ([]string, error) {
	return s.client.HKeys(ctx, roomUsersKey(roomID)).Result()
}

func (s *roomPresenceStore) IsConnectedUser(ctx context.Context, roomID, userID string) (bool, error) {
	return s.client.HExists(ctx, roomUsersKey(roomID), userID).Result()
}

func (s *roomPresenceStore) CountConnectedUsers(ctx context.Context, roomID string) (int64, error) {
	return s.client.HLen(ctx, roomUsersKey(roomID)).Result()
}

func (s *roomPresenceStore) DeleteRoomConnections(ctx context.Context, roomID string) error {
	return s.client.Del(ctx, roomUsersKey(roomID)).Err()
}

func roomUsersKey(roomID string) string {
	return "room:" + roomID + ":users"
}
