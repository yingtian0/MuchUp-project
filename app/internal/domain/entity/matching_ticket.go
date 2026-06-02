package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type MatchingTicketID string

type MatchingTicketStatus string

const (
	MatchingWaiting   MatchingTicketStatus = "WAITING"
	MatchingMatched   MatchingTicketStatus = "MATCHED"
	MatchingCancelled MatchingTicketStatus = "CANCELLED"
	MatchingExpired   MatchingTicketStatus = "EXPIRED"
)

type MatchingTicket struct {
	ID             MatchingTicketID
	UserID         UserID
	Status         MatchingTicketStatus
	IdempotencyKey string

	RequestedAt   time.Time
	MatchedRoomID *RoomID
	ExpiresAt     time.Time
}

func NewMatchingTicket(userID UserID, idempotencyKey string, now time.Time) (*MatchingTicket, error) {
	if userID == "" {
		return nil, errors.New("user_id is required")
	}

	if idempotencyKey == "" {
		return nil, errors.New("idempotency_key is required")
	}

	return &MatchingTicket{
		ID:             MatchingTicketID(uuid.NewString()),
		UserID:         userID,
		Status:         MatchingWaiting,
		IdempotencyKey: idempotencyKey,
		RequestedAt:    now,
		ExpiresAt:      now.Add(5 * time.Minute),
	}, nil
}
