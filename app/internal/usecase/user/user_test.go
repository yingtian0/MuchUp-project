package user

import (
	"errors"
	"testing"

	usecase "MuchUp/app/internal/controllers/usecase"
)

func TestUnimplementedMethodsReturnErrNotImplemented(t *testing.T) {
	u := &userUsecase{}

	_, err := u.Login("", "")
	if !errors.Is(err, usecase.ErrNotImplemented) {
		t.Errorf("Login() error = %v, want ErrNotImplemented", err)
	}

	if err := u.JoinRoom("", ""); !errors.Is(err, usecase.ErrNotImplemented) {
		t.Errorf("JoinRoom() error = %v, want ErrNotImplemented", err)
	}

	if err := u.LeaveRoom("", ""); !errors.Is(err, usecase.ErrNotImplemented) {
		t.Errorf("LeaveRoom() error = %v, want ErrNotImplemented", err)
	}
}
