package v1

import (
	"context"
	"time"

	"MuchUp/app/internal/domain/entity"

	authv1 "MuchUp/app/proto/gen/go/auth/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (h *GrpcHandler) SignUp(ctx context.Context, req *authv1.SignUpRequest) (*authv1.SignUpResponse, error) {
	user := &entity.User{
		NickName:     req.GetName(),
		Email:        &req.Email,
		PasswordHash: req.GetPassword(),
	}

	createdUser, err := h.userUsecase.CreateUser(user)
	if err != nil {
		return nil, h.handleError(ctx, "SignUp", err)
	}

	email := ""
	if createdUser.Email != nil {
		email = *createdUser.Email
	}

	return &authv1.SignUpResponse{
		UserId:    createdUser.ID,
		Name:      createdUser.NickName,
		Email:     email,
		CreatedAt: timestamppb.New(createdUser.CreatedAt),
	}, nil
}

func (h *GrpcHandler) Login(ctx context.Context, req *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	token, err := h.userUsecase.Login(req.GetEmail(), req.GetPassword())
	if err != nil {
		return nil, h.handleError(ctx, "Login", err)
	}

	expiresAt := time.Now().Add(24 * time.Hour)
	return &authv1.LoginResponse{
		AccessToken:  token,
		RefreshToken: "",
		ExpiresAt:    timestamppb.New(expiresAt),
	}, nil
}

func (h *GrpcHandler) UpdateUser(ctx context.Context, req *authv1.UpdateUserRequest) (*authv1.UpdateUserResponse, error) {
	user, err := h.userUsecase.GetUserByID(req.GetUserId())
	if err != nil {
		return nil, h.handleError(ctx, "UpdateUser.GetUserByID", err)
	}

	if req.GetName() != "" {
		user.NickName = req.GetName()
	}

	updatedUser, err := h.userUsecase.UpdateUser(user)
	if err != nil {
		return nil, h.handleError(ctx, "UpdateUser", err)
	}

	return &authv1.UpdateUserResponse{
		UserId: updatedUser.ID,
		Name:   updatedUser.NickName,
	}, nil
}

func (h *GrpcHandler) DeleteUser(ctx context.Context, req *authv1.DeleteUserRequest) (*authv1.DeleteUserResponse, error) {
	if err := h.userUsecase.DeleteUser(req.GetUserId()); err != nil {
		return nil, h.handleError(ctx, "DeleteUser", err)
	}

	return &authv1.DeleteUserResponse{
		UserId: req.GetUserId(),
	}, nil
}
