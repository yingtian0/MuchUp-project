package entity

import "time"

type IcebreakPromptType string

const (
	IcebreakPromptInitialTopic    IcebreakPromptType = "INITIAL_TOPIC"
	IcebreakPromptFollowUp        IcebreakPromptType = "FOLLOW_UP"
	IcebreakPromptSilenceRecovery IcebreakPromptType = "SILENCE_RECOVERY"
	IcebreakPromptMission         IcebreakPromptType = "MISSION"
	IcebreakPromptSummary         IcebreakPromptType = "SUMMARY"
)

type IcebreakTriggerType string

const (
	IcebreakTriggerRoomActivated IcebreakTriggerType = "ROOM_ACTIVATED"
	IcebreakTriggerUserMessage   IcebreakTriggerType = "USER_MESSAGE"
	IcebreakTriggerSilence       IcebreakTriggerType = "SILENCE"
	IcebreakTriggerManual        IcebreakTriggerType = "MANUAL"
)

type IcebreakPrompt struct {
	ID               string
	RoomID           RoomID
	MessageID        *MessageID
	Type             IcebreakPromptType
	Trigger          IcebreakTriggerType
	TargetUserID     *UserID
	SourceMessageIDs []MessageID
	GeneratedText    string
	CreatedAt        time.Time
}
