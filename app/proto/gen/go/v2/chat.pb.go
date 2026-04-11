package v2

import timestamppb "google.golang.org/protobuf/types/known/timestamppb"

type User struct {
	Id        string
	NickName  string
	Email     string
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Message struct {
	MessageId string
	SenderId  string
	GroupId   string
	Text      string
	CreatedAt *timestamppb.Timestamp
}

type CreateUserRequest struct {
	NickName string
	Email    string
	Password string
}

func (x *CreateUserRequest) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUserRequest struct {
	Id string
}

func (x *GetUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateMessageRequest struct {
	SenderId string
	GroupId  string
	Text     string
}

func (x *CreateMessageRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *CreateMessageRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *CreateMessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetMessagesByGroupRequest struct {
	GroupId string
	Limit   int32
	Offset  int32
}

func (x *GetMessagesByGroupRequest) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GetMessagesByGroupRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetMessagesByGroupRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}
