package entity

import (
	"errors"
	"time"
)

type UserID string

type UserStatus string
type PrimaryAuthMethod string

const (
	UserStatusActive    UserStatus = "ACTIVE"
	UserStatusSuspended UserStatus = "SUSPENDED"
	UserStatusDeleted   UserStatus = "DELETED"
)

const (
	AuthMethodEmail PrimaryAuthMethod = "email"
	AuthMethodPhone PrimaryAuthMethod = "phone"
)

type UserProfile struct {
	ID            UserID
	NickName      string
	DisplayName   string
	Email         *string
	PhoneNumber   *string
	PasswordHash  string
	EmailVerified bool
	PhoneVerified bool
	AuthMethod    PrimaryAuthMethod
	AvatarURL     *string
	Interests     []string
	Hobbies       []string
	FavoriteTags  []string
	UsagePurpose  string
	IsActive      bool
	IsBanned      bool

	Status UserStatus
	IsBot  bool

	BlockedUserIDs map[UserID]struct{}

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserProfile) CanJoinMatching() bool {
	return u.Status == UserStatusActive && !u.IsBot && u.IsActive && !u.IsBanned
}

func (u *UserProfile) CanSendMessageTo(target UserID) bool {
	if u.Status != UserStatusActive {
		return false
	}

	if _, blocked := u.BlockedUserIDs[target]; blocked {
		return false
	}

	return true
}

func (u *UserProfile) Block(target UserID) error {
	if u.ID == target {
		return errors.New("cannot block yourself")
	}

	if u.BlockedUserIDs == nil {
		u.BlockedUserIDs = map[UserID]struct{}{}
	}

	u.BlockedUserIDs[target] = struct{}{}

	return nil
}
