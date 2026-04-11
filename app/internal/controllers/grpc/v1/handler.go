package v1

import (
	"MuchUp/backend/internal/domain/repository"
	"MuchUp/backend/internal/domain/usecase"
	"MuchUp/backend/pkg/logger"

	authv1 "MuchUp/backend/proto/gen/go/auth/v1"
	chatv1 "MuchUp/backend/proto/gen/go/chat/v1"
)

type GrpcHandler struct {
	authv1.UnimplementedAuthServiceServer
	chatv1.UnimplementedChatServiceServer

	userUsecase    usecase.UserUsecase
	messageUsecase usecase.MessageUsecase
	groupRepo      repository.ChatGroupRepository
	logger         logger.Logger
}

func NewGrpcHandler(
	userUsecase usecase.UserUsecase,
	messageUsecase usecase.MessageUsecase,
	groupRepo repository.ChatGroupRepository,
	logger logger.Logger,
) *GrpcHandler {
	return &GrpcHandler{
		userUsecase:    userUsecase,
		messageUsecase: messageUsecase,
		groupRepo:      groupRepo,
		logger:         logger,
	}
}
