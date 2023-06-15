package service

import (
	api "github.com/mephistolie/chefbook-backend-tag/api/proto/implementation/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Tag struct {
	api.TagServiceClient
	Conn *grpc.ClientConn
}

func NewTag(addr string) (*Tag, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		return nil, err
	}
	return &Tag{
		TagServiceClient: api.NewTagServiceClient(conn),
		Conn:             conn,
	}, nil
}
