package entity

import "time"

type RoomType string
type RoomStatus string

const (
	RoomTypeChat RoomType = "chat"

	RoomWaiting RoomStatus = "waiting"
	RoomActive  RoomStatus = "active"
	RoomClosed  RoomStatus = "closed"
)

type Room struct {
	ID        RoomID                 `json:"id"`
	Type      RoomType               `json:"type"`
	Status    RoomStatus             `json:"status"`
	Capacity  int                    `json:"capacity"`
	Members   map[UserID]*RoomMember `json:"members"`
	CreatedBy UserID                 `json:"created_by"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
}
