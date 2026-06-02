package entity

import "time"

type IcebreakSessionStatus string

const (
	IcebreakSessionActive    IcebreakSessionStatus = "ACTIVE"
	IcebreakSessionCompleted IcebreakSessionStatus = "COMPLETED"
	IcebreakSessionCancelled IcebreakSessionStatus = "CANCELLED"
)

type IcebreakSession struct {
	ID                   string
	RoomID               RoomID
	Status               IcebreakSessionStatus
	StartedAt            time.Time
	LastAIMessageAt      *time.Time
	LastUserMessageAt    *time.Time
	InterventionCount    int
	MaxInterventionCount int
	CooldownSeconds      int
	InitialTopicSent     bool
}
