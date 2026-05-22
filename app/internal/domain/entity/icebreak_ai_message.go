package entity

import "time"

type IcebreakAIMessageType string

const (
	IcebreakAIMessageInitialTopic    IcebreakAIMessageType = "INITIAL_TOPIC"
	IcebreakAIMessageFollowUp        IcebreakAIMessageType = "FOLLOW_UP"
	IcebreakAIMessageSilenceRecovery IcebreakAIMessageType = "SILENCE_RECOVERY"
	IcebreakAIMessageMission         IcebreakAIMessageType = "MISSION"
	IcebreakAIMessageSummary         IcebreakAIMessageType = "SUMMARY"
)

type IcebreakTriggerType string

const (
	IcebreakTriggerRoomActivated IcebreakTriggerType = "ROOM_ACTIVATED"
	IcebreakTriggerUserMessage   IcebreakTriggerType = "USER_MESSAGE"
	IcebreakTriggerSilence       IcebreakTriggerType = "SILENCE"
	IcebreakTriggerManual        IcebreakTriggerType = "MANUAL"
)

type IcebreakAIMessage struct {
	ID               string
	RoomID           RoomID
	MessageID        *MessageID
	Type             IcebreakAIMessageType
	Trigger          IcebreakTriggerType
	TargetUserID     *UserID
	SourceMessageIDs []MessageID
	GeneratedText    string
	CreatedAt        time.Time
}
