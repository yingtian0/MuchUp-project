package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type RoomID string

type RoomStatus string

const (
	RoomWaiting RoomStatus = "WAITING"
	RoomActive  RoomStatus = "ACTIVE"
	RoomClosed  RoomStatus = "CLOSED"
	RoomExpired RoomStatus = "EXPIRED"
)

type RoomType string

const (
	RoomTypeRandomGroup RoomType = "RANDOM_GROUP"
)

const RandomRoomCapacity = 5

type Room struct {
	ID       RoomID
	Type     RoomType
	Status   RoomStatus
	Capacity int

	Members map[UserID]*RoomMember

	CreatedBy   UserID
	CreatedAt   time.Time
	ActivatedAt *time.Time
	ClosedAt    *time.Time

	Version int64
}

func NewRandomRoom(owner UserID, now time.Time) (*Room, error) {
	if owner == "" {
		return nil, errors.New("owner is required")
	}

	room := &Room{
		ID:        RoomID(uuid.NewString()),
		Type:      RoomTypeRandomGroup,
		Status:    RoomWaiting,
		Capacity:  RandomRoomCapacity,
		Members:   map[UserID]*RoomMember{},
		CreatedBy: owner,
		CreatedAt: now,
		Version:   1,
	}

	if err := room.Join(owner, RoomMemberRoleOwner, now); err != nil {
		return nil, err
	}
	return room, nil
}

func (r *Room) Join(userID UserID, role RoomMemberRole, now time.Time) error {
	if userID == "" {
		return errors.New("user_id is required")
	}
	if r.Status != RoomWaiting {
		return errors.New("room is not joinable")
	}
	if _, exists := r.Members[userID]; exists {
		return nil
	}
	if len(r.Members) >= r.Capacity {
		return errors.New("room is full")
	}

	r.Members[userID] = &RoomMember{
		UserID:   userID,
		Role:     role,
		Status:   RoomMemberJoined,
		JoinedAt: now,
	}

	if len(r.Members) == r.Capacity {
		r.Status = RoomActive
		r.ActivatedAt = &now
	}

	r.Version++
	return nil
}

func (r *Room) CanSendMessage(userID UserID) bool {
	member, ok := r.Members[userID]
	return ok && r.Status == RoomActive && member.Status == RoomMemberJoined
}
