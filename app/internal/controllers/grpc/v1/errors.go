package v1

import (
	"context"
	"errors"
	"fmt"

	"MuchUp/app/internal/domain/usecase"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (h *GrpcHandler) handleError(ctx context.Context, operation string, err error) error {
	if err == nil {
		return nil
	}

	h.logger.WithContext(ctx).WithError(err).Errorf("Failed to %s", operation)

	switch {
	case errors.Is(err, usecase.ErrNotFound), errors.Is(err, gorm.ErrRecordNotFound):
		return status.Error(codes.NotFound, fmt.Sprintf("%s not found", operation))
	case errors.Is(err, usecase.ErrInvalidArgument):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, usecase.ErrPermissionDenied):
		return status.Error(codes.PermissionDenied, "permission denied")
	default:
		return status.Error(codes.Internal, fmt.Sprintf("failed to %s", operation))
	}
}
