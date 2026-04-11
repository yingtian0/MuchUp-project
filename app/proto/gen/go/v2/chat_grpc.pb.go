package v2

import (
	context "context"

	grpc "google.golang.org/grpc"
)

type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
}

type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*User, error) {
	return nil, nil
}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, nil
}

func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                    {}

type MessageServiceServer interface {
	CreateMessage(context.Context, *CreateMessageRequest) (*Message, error)
	GetMessagesByGroup(*GetMessagesByGroupRequest, MessageService_GetMessagesByGroupServer) error
	mustEmbedUnimplementedMessageServiceServer()
}

type UnimplementedMessageServiceServer struct{}

func (UnimplementedMessageServiceServer) CreateMessage(context.Context, *CreateMessageRequest) (*Message, error) {
	return nil, nil
}

func (UnimplementedMessageServiceServer) GetMessagesByGroup(*GetMessagesByGroupRequest, MessageService_GetMessagesByGroupServer) error {
	return nil
}

func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}
func (UnimplementedMessageServiceServer) testEmbeddedByValue()                         {}

type MessageService_GetMessagesByGroupServer interface {
	Send(*Message) error
	grpc.ServerStream
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {}
