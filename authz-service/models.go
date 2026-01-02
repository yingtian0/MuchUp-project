package authzservice

import "time"


type RoomState struct {
	RoomID string `json:"room_id"`
	OwnerID string `json:"owner_id"`
	IsFull bool `json:"is_full"`
	CreatedAt time.Time `json:"created_at"`
}