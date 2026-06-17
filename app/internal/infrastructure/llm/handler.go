package llm

import (
	"context"
	"time"

	usecase "MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/store"
	llmv1 "MuchUp/app/proto/gen/go/llm/v1"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

const aiAgentUserID = "ai_agent"

type Handler struct {
	client        llmv1.LLMServiceClient
	messageStream store.MessageStreamStore
}

var _ usecase.LLMHandler = (*Handler)(nil)

func NewHandler(conn *grpc.ClientConn, messageStream store.MessageStreamStore) *Handler {
	return &Handler{
		client:        llmv1.NewLLMServiceClient(conn),
		messageStream: messageStream,
	}
}

func (h *Handler) HandleRoomCreated(ctx context.Context, room *entity.Room, owner *entity.UserProfile) error {
	ownerName := owner.NickName
	if ownerName == "" {
		ownerName = owner.DisplayName
	}

	request := &llmv1.GenerateReplyRequest{
		RoomId:       string(room.ID),
		SessionId:    string(room.ID),
		TargetUserId: string(owner.ID),
		SystemPrompt: "You are a chat facilitator for a newly created social room. Reply in Japanese with one short welcoming message.",
		Model:        "facilitator-v1",
		Temperature:  0.7,
		MaxTokens:    120,
		Messages: []*llmv1.ContextMessage{
			{
				MessageId: "room-created",
				RoomId:    string(room.ID),
				UserId:    string(owner.ID),
				Role:      "system",
				Content:   "room created for " + ownerName,
				CreatedAt: time.Now().Unix(),
			},
		},
		Metadata: map[string]string{
			"event":      "room_created",
			"owner_name": ownerName,
		},
	}

	response, err := h.client.GenerateReply(ctx, request)
	if err != nil {
		return err
	}

	if response.GetContent() == "" {
		return nil
	}

	createdAt := time.Unix(response.GetCreatedAt(), 0)
	if response.GetCreatedAt() == 0 {
		createdAt = time.Now()
	}

	content := response.GetContent()
	message := &entity.Message{
		ID:              entity.MessageID(uuid.NewString()),
		ClientMessageID: new("room-created"),
		SenderID:        entity.UserID(aiAgentUserID),
		RoomID:          room.ID,
		Text:            &content,
		CreatedAt:       createdAt,
		UpdatedAt:       createdAt,
	}

	_, err = h.messageStream.AppendMessage(ctx, message)

	return err
}
