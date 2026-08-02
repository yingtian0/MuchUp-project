package user

import (
	"context"
	"errors"

	usecase "MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

// TODO: 他依存の注入が増える想定のため、ユースケース初期化をここに集約する
func NewUserUsecase(
	userRepo repository.UserRepository,
) usecase.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
func (u *userUsecase) CreateUser(user *entity.UserProfile) (*entity.UserProfile, error) {
	ctx := context.Background()

	if user == nil {
		return nil, errors.New("user is required")
	}

	if err := u.userRepo.Insert(ctx, user); err != nil {
		return nil, err
	}

	created, err := u.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (u *userUsecase) GetUserByID(id string) (*entity.UserProfile, error) {
	return u.userRepo.FindByID(context.Background(), entity.UserID(id))
}

func (u *userUsecase) GetUserByEmail(email string) (*entity.UserProfile, error) {
	return u.userRepo.FindByEmail(context.Background(), email)
}

func (u *userUsecase) UpdateUser(user *entity.UserProfile) (*entity.UserProfile, error) {
	ctx := context.Background()

	if user == nil {
		return nil, errors.New("user is required")
	}

	if err := u.userRepo.Update(ctx, user); err != nil {
		return nil, err
	}

	updated, err := u.userRepo.FindByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.userRepo.Delete(context.Background(), entity.UserID(id))
}

func (u *userUsecase) GetUsers(limit, offset uint) ([]*entity.UserProfile, error) {
	return u.userRepo.FindAll(context.Background(), limit, offset)
}

// TODO: 認証基盤の実装に合わせてログイン処理を実装する
func (u *userUsecase) Login(_, _ string) (string, error) {
	return "", usecase.ErrNotImplemented
}

// TODO: ルーム所属管理の実装に合わせて参加処理を実装する
func (u *userUsecase) JoinRoom(_, _ string) error {
	return usecase.ErrNotImplemented
}

// TODO: ルーム所属管理の実装に合わせて退出処理を実装する
func (u *userUsecase) LeaveRoom(_, _ string) error {
	return usecase.ErrNotImplemented
}

func (u *userUsecase) GetUsersByRoom(roomID string) ([]*entity.UserProfile, error) {
	return u.userRepo.FindByRoom(context.Background(), entity.RoomID(roomID))
}
