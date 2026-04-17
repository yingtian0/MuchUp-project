package v1

import (
	"context"

	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/usecase"

	chatv1 "MuchUp/app/proto/gen/go/chat/v1"
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
	message, err := entity.NewMessage(req.GetUserId(), req.GetRoomId(), req.GetText())
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
	groups, err := h.groupRepo.GetGroupByUserID(req.GetUserId())
	if err != nil {
		return nil, h.handleError(ctx, "MatchRoom.GetGroupByUserID", err)
	}
	if len(groups) == 0 {
		return nil, h.handleError(ctx, "MatchRoom.GetGroupByUserID", usecase.ErrNotFound)
	}

	group := groups[0]
	ownerID := ""
	if len(group.Members) > 0 {
		ownerID = group.Members[0].ID
	}

	return &chatv1.MatchRoomResponse{
		OwnerId: ownerID,
		RooomId: group.ID,
	}, nil
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
