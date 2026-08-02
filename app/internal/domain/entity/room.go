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
	Capacity           uint
	Members            map[UserID]*RoomMember
	CreatedBy          *UserID
	CreatedAt          time.Time
	ActivatedAt        *time.Time
	ClosedAt           *time.Time
	LastMessageAt      *time.Time
	LastAIIntervenedAt *time.Time
}

// ReconstructRoom は永続化済みのデータから Room を復元する。
// New系と異なりバリデーションは行わない。Members は別途取得して設定する。
func ReconstructRoom(
	id RoomID,
	roomType RoomType,
	status RoomStatus,
	capacity uint,
	createdBy *UserID,
	createdAt time.Time,
	activatedAt *time.Time,
	closedAt *time.Time,
	lastMessageAt *time.Time,
	lastAIIntervenedAt *time.Time,
) *Room {
	return &Room{
		ID:                 id,
		Type:               roomType,
		Status:             status,
		Capacity:           capacity,
		CreatedBy:          createdBy,
		CreatedAt:          createdAt,
		ActivatedAt:        activatedAt,
		ClosedAt:           closedAt,
		LastMessageAt:      lastMessageAt,
		LastAIIntervenedAt: lastAIIntervenedAt,
	}
}
