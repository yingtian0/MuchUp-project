package v1

import (
	"MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/repository"
	"MuchUp/app/pkg/logger"

	authv1 "MuchUp/app/proto/gen/go/auth/v1"
	chatv1 "MuchUp/app/proto/gen/go/chat/v1"
)

type GrpcHandler struct {
	authv1.UnimplementedAuthServiceServer
	chatv1.UnimplementedChatServiceServer

	userUsecase    usecase.UserUsecase
	messageUsecase usecase.MessageUsecase
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
		logger:         logger,
	}
}
