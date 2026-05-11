package entity

import (
	"errors"
	"time"
)

type UserID string

type UserStatus string

const (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"
	UserStatusDeleted   UserStatus = "DELETED"
)

type UserProfile struct {
	ID          UserID
	DisplayName string
	AvatarURL   *string

	Status UserStatus
	IsBot  bool

	BlockedUserIDs map[UserID]bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserProfile) CanJoinMatching() bool {
	return u.Status == UserStatusActive && !u.IsBot
}

func (u *UserProfile) CanSendMessageTo(target UserID) bool {
	if u.Status != UserStatusActive {
		return false
	}
	return !u.BlockedUserIDs[target]
}

func (u *UserProfile) Block(target UserID) error {
	if u.ID == target {
		return errors.New("cannot block yourself")
	}
	if u.BlockedUserIDs == nil {
		u.BlockedUserIDs = map[UserID]bool{}
	}
	u.BlockedUserIDs[target] = true
	return nil
}
