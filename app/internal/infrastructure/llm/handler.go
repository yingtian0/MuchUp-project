package llm

import (
	usecase "MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	llmv1 "MuchUp/app/proto/gen/go/llm/v1"
	"MuchUp/app/utils"
	"context"
	"time"

	"google.golang.org/grpc"
)

const aiAgentUserID = "ai_agent"

type Handler struct {
	client        llmv1.LLMServiceClient
	messageStream repository.MessageStreamStore
}

var _ usecase.LLMHandler = (*Handler)(nil)

func NewHandler(conn *grpc.ClientConn, messageStream repository.MessageStreamStore) *Handler {
	return &Handler{
		client:        llmv1.NewLLMServiceClient(conn),
		messageStream: messageStream,
	}
}

func (h *Handler) HandleRoomCreated(ctx context.Context, group *entity.ChatGroup, owner *entity.User) error {
	ownerName := owner.NickName
	request := &llmv1.GenerateReplyRequest{
		RoomId:       group.ID,
		SessionId:    group.ID,
		TargetUserId: owner.ID,
		SystemPrompt: "You are a chat facilitator for a newly created social room. Reply in Japanese with one short welcoming message.",
		Model:        "facilitator-v1",
		Temperature:  0.7,
		MaxTokens:    120,
		Messages: []*llmv1.ContextMessage{
			{
				MessageId: "room-created",
				RoomId:    group.ID,
				UserId:    owner.ID,
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
		MessageID: utils.GenerateUUID(),
		SenderID:  aiAgentUserID,
		GroupID:   group.ID,
		Text:      &content,
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}
	_, err = h.messageStream.AppendMessage(ctx, message)
	return err
}
