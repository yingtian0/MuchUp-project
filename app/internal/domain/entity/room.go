package entity

import "time"

type RoomType string
type RoomStatus string

const (
	RoomTypeClosed RoomType = "Closed"
	RoomTypeRandom RoomType = "random"

	RoomWaiting RoomStatus = "waiting"
	RoomActive  RoomStatus = "active"
	RoomClosed  RoomStatus = "closed"
)

type Room struct {
	ID                 RoomID
	Type               RoomType
	Status             RoomStatus
	Capacity           int
	Members            map[UserID]*RoomMember
	CreatedBy          *UserID
	CreatedAt          time.Time
	ActivatedAt        *time.Time
	ClosedAt           *time.Time
	LastMessageAt      *time.Time
	LastAIIntervenedAt *time.Time
}
