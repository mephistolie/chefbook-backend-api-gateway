package service

import (
	api "github.com/mephistolie/chefbook-backend-user/api/proto/implementation/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	api.UserServiceClient
	Conn *grpc.ClientConn
}

func NewUser(addr string) (*User, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		return nil, err
	}
	return &User{
		UserServiceClient: api.NewUserServiceClient(conn),
		Conn:              conn,
	}, nil
}
