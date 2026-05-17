package user

import (
	usecase "MuchUp/app/internal/controllers/usecase"
	"MuchUp/app/internal/domain/entity"
	"MuchUp/app/internal/domain/repository"
	"errors"
)

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(
	userRepo repository.UserRepository,
) usecase.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
func (u *userUsecase) CreateUser(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, errors.New("user is required")
	}
	if err := u.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	created, err := u.userRepo.GetUserByID(user.ID)
	if err != nil {
		return user, nil
	}
	return created, nil
}

func (u *userUsecase) GetUserByID(id string) (*entity.User, error) {
	return u.userRepo.GetUserByID(id)
}

func (u *userUsecase) GetUserByEmail(email string) (*entity.User, error) {
	return u.userRepo.GetUserByEmail(email)
}

func (u *userUsecase) UpdateUser(user *entity.User) (*entity.User, error) {
	if user == nil {
		return nil, errors.New("user is required")
	}
	if err := u.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}
	updated, err := u.userRepo.GetUserByID(user.ID)
	if err != nil {
		return user, nil
	}
	return updated, nil
}

func (u *userUsecase) DeleteUser(id string) error {
	return u.userRepo.DeleteUser(id)
}

func (u *userUsecase) GetUsers(limit, offset int) ([]*entity.User, error) {
	return u.userRepo.GetUsers(limit, offset)
}

func (u *userUsecase) Login(email, password string) (string, error) {
	return "", errors.New("not implemented")
}

func (u *userUsecase) JoinRoom(userID, roomID string) error {
	return errors.New("not implemented")
}

func (u *userUsecase) LeaveRoom(userID, roomID string) error {
	return errors.New("not implemented")
}

func (u *userUsecase) GetUsersByRoom(roomID string) ([]*entity.User, error) {
	return u.userRepo.GetUsersByRoom(roomID)
}
