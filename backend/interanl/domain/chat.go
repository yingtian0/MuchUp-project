package domain

type Message struct {
	MessageID string
	SenderID  string
	RoomID    string
	Text      string
}

type Room struct {
	roomId string 
	isFull bool
	Members []User
}