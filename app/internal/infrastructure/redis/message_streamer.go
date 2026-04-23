package redis

import (
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"context"
	"fmt"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type MessageStreamStore struct {
	client goredis.Cmdable
	maxLen int64
}

var _ repository.MessageStreamStore = (*MessageStreamStore)(nil)

func NewMessageStreamStore(client goredis.Cmdable, maxLen int64) *MessageStreamStore {
	return &MessageStreamStore{
		client: client,
		maxLen: maxLen,
	}
}

func (s *MessageStreamStore) AppendMessage(ctx context.Context, message *entity.Message) (string, error) {
	args := &goredis.XAddArgs{
		Stream: streamKey(message.GroupID),
		Values: map[string]any{
			"message_id": message.MessageID,
			"sender_id":  message.SenderID,
			"group_id":   message.GroupID,
			"text":       derefString(message.Text),
			"image":      derefString(message.Image),
			"created_at": strconv.FormatInt(message.CreatedAt.UnixMilli(), 10),
			"updated_at": strconv.FormatInt(message.UpdatedAt.UnixMilli(), 10),
		},
	}
	if s.maxLen > 0 {
		args.MaxLen = s.maxLen
		args.Approx = true
	}
	return s.client.XAdd(ctx, args).Result()
}

func (s *MessageStreamStore) GetRecentMessages(ctx context.Context, roomID string, count int64) ([]*entity.Message, error) {
	streams, err := s.client.XRevRangeN(ctx, streamKey(roomID), "+", "-", count).Result()
	if err != nil {
		return nil, err
	}

	messages := make([]*entity.Message, 0, len(streams))
	for i := len(streams) - 1; i >= 0; i-- {
		message, err := streamMessageToEntity(streams[i])
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (s *MessageStreamStore) GetMessagesAfter(ctx context.Context, roomID, lastMessageID string, count int64) ([]*entity.Message, error) {
	start := "-"
	if lastMessageID != "" {
		start = "(" + lastMessageID
	}

	streams, err := s.client.XRangeN(ctx, streamKey(roomID), start, "+", count).Result()
	if err != nil {
		return nil, err
	}

	messages := make([]*entity.Message, 0, len(streams))
	for _, stream := range streams {
		message, err := streamMessageToEntity(stream)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (s *MessageStreamStore) DeleteMessageHistory(ctx context.Context, roomID string) error {
	return s.client.Del(ctx, streamKey(roomID)).Err()
}

func streamMessageToEntity(stream goredis.XMessage) (*entity.Message, error) {
	createdAt, err := parseUnixMilli(stream.Values["created_at"])
	if err != nil {
		return nil, err
	}
	updatedAt, err := parseUnixMilli(stream.Values["updated_at"])
	if err != nil {
		return nil, err
	}

	message := &entity.Message{
		MessageID: toString(stream.Values["message_id"]),
		SenderID:  toString(stream.Values["sender_id"]),
		GroupID:   toString(stream.Values["group_id"]),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if text := toString(stream.Values["text"]); text != "" {
		message.Text = &text
	}
	if image := toString(stream.Values["image"]); image != "" {
		message.Image = &image
	}
	if video := toString(stream.Values["video"]); video != "" {
		message.Video = &video
	}
	if sticker := toString(stream.Values["sticker"]); sticker != "" {
		message.Sticker = &sticker
	}

	return message, nil
}

func parseUnixMilli(value any) (time.Time, error) {
	if value == nil || toString(value) == "" {
		return time.Time{}, nil
	}

	unixMilli, err := strconv.ParseInt(toString(value), 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.UnixMilli(unixMilli), nil
}

func derefString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func toString(value any) string {
	switch typed := value.(type) {
	case nil:
		return ""
	case string:
		return typed
	case []byte:
		return string(typed)
	default:
		return fmt.Sprint(typed)
	}
}

func streamKey(roomID string) string {
	return "room:" + roomID + ":messages"
}
