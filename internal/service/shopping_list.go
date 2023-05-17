package service

import (
	api "github.com/mephistolie/chefbook-backend-shopping-list/api/v2/proto/implementation/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ShoppingList struct {
	api.ShoppingListServiceClient
	Conn *grpc.ClientConn
}

func NewShoppingList(addr string) (*ShoppingList, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		return nil, err
	}
	return &ShoppingList{
		ShoppingListServiceClient: api.NewShoppingListServiceClient(conn),
		Conn:                      conn,
	}, nil
}
