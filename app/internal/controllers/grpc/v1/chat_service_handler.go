package v1

import (
	"context"

	"MuchUp/app/internal/domain/entity"

	chatv1 "MuchUp/app/proto/gen/go/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *GrpcHandler) GetMessages(ctx context.Context, req *chatv1.GetMessagesRequest) (*chatv1.GetMessagesResponse, error) {
	messages, err := h.messageUsecase.GetMessagesByGroup(req.GetRoomId(), 50, 0)
	if err != nil {
		return nil, h.handleError(ctx, "GetMessages", err)
	}

	res := &chatv1.GetMessagesResponse{
		Message: make([]*chatv1.Message, 0, len(messages)),
	}
	for _, msg := range messages {
		res.Message = append(res.Message, toChatMessage(msg))
	}

	return res, nil
}

func (h *GrpcHandler) SendMessage(ctx context.Context, req *chatv1.SendMessageRequest) (*chatv1.SendMessageResponse, error) {
	if req.GetMessage() == nil {
		return nil, status.Error(codes.InvalidArgument, "message is required")
	}

	message, err := entity.NewMessage(req.GetMessage().GetSenderId(), req.GetRoomId(), req.GetMessage().GetText())
	if err != nil {
		return nil, h.handleError(ctx, "SendMessage.NewMessage", err)
	}

	createdMessage, err := h.messageUsecase.CreateMessage(message)
	if err != nil {
		return nil, h.handleError(ctx, "SendMessage.CreateMessage", err)
	}

	return &chatv1.SendMessageResponse{
		RoomId:    createdMessage.GroupID,
		MessageId: createdMessage.MessageID,
		CreatedAt: createdMessage.CreatedAt.Unix(),
	}, nil
}

func (h *GrpcHandler) MatchRoom(ctx context.Context, req *chatv1.MatchRoomRequest) (*chatv1.MatchRoomResponse, error) {
	return nil, status.Error(codes.Unimplemented, "MatchRoom requires a user identifier source outside the request body")
}

func toChatMessage(msg *entity.Message) *chatv1.Message {
	if msg == nil {
		return nil
	}

	text := ""
	if msg.Text != nil {
		text = *msg.Text
	}

	return &chatv1.Message{
		MessageId: msg.MessageID,
		SenderId:  msg.SenderID,
		Text:      text,
		CreatedAt: msg.CreatedAt.Unix(),
	}
}
