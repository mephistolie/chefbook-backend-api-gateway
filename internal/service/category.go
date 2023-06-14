package service

import (
	api "github.com/mephistolie/chefbook-backend-category/api/proto/implementation/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Category struct {
	api.CategoryServiceClient
	Conn *grpc.ClientConn
}

func NewCategory(addr string) (*Category, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		return nil, err
	}
	return &Category{
		CategoryServiceClient: api.NewCategoryServiceClient(conn),
		Conn:                  conn,
	}, nil
}
