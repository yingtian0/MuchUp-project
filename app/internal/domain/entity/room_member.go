package entity

import "time"

type RoomMemberRole string

const (
	RoomMemberRoleOwner  RoomMemberRole = "OWNER"
	RoomMemberRoleMember RoomMemberRole = "MEMBER"
	RoomMemberRoleAI     RoomMemberRole = "AI"
)

type RoomMemberStatus string

const (
	RoomMemberJoined RoomMemberStatus = "JOINED"
	RoomMemberLeft   RoomMemberStatus = "LEFT"
	RoomMemberKicked RoomMemberStatus = "KICKED"
)

type RoomMember struct {
	UserID UserID
	Role   RoomMemberRole
	Status RoomMemberStatus

	JoinedAt time.Time
	LeftAt   *time.Time

	LastReadMessageID *MessageID
}
