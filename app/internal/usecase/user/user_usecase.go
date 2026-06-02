package user

import (
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
	if user == nil {
		return nil, errors.New("user is required")
	}

	if err := u.userRepo.CreateUser(user); err != nil {
		return nil, err
	}

	created, err := u.userRepo.GetUserByID(string(user.ID))
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (u *userUsecase) GetUserByID(id string) (*entity.UserProfile, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUsecase) GetUserByEmail(email string) (*entity.UserProfile, error) {
	return u.userRepo.GetUserByEmail(email)
}

func (u *userUsecase) UpdateUser(user *entity.UserProfile) (*entity.UserProfile, error) {
	if user == nil {
		return nil, errors.New("user is required")
	}

	if err := u.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}

	updated, err := u.userRepo.GetUserByID(string(user.ID))
	if err != nil {
		return nil, err
	}

	return updated, nil
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.userRepo.DeleteUser(id)
}

func (u *userUsecase) GetUsers(limit, offset int) ([]*entity.UserProfile, error) {
	return u.userRepo.GetUsers(limit, offset)
}

// TODO: 認証基盤の実装に合わせてログイン処理を実装する
func (u *userUsecase) Login(_, _ string) (string, error) {
	return "", errors.New("not implemented")
}

// TODO: ルーム所属管理の実装に合わせて参加処理を実装する
func (u *userUsecase) JoinRoom(_, _ string) error {
	return errors.New("not implemented")
}

// TODO: ルーム所属管理の実装に合わせて退出処理を実装する
func (u *userUsecase) LeaveRoom(_, _ string) error {
	return errors.New("not implemented")
}

func (u *userUsecase) GetUsersByRoom(roomID string) ([]*entity.UserProfile, error) {
	return u.userRepo.GetUsersByRoom(roomID)
}
