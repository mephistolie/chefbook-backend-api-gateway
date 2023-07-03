package service

import (
	api "github.com/mephistolie/chefbook-backend-encryption/api/proto/implementation/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Encryption struct {
	api.EncryptionServiceClient
	Conn *grpc.ClientConn
}

func NewEncryption(addr string) (*Encryption, error) {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.Dial(addr, opts)
	if err != nil {
		return nil, err
	}
	return &Encryption{
		EncryptionServiceClient: api.NewEncryptionServiceClient(conn),
		Conn:                    conn,
	}, nil
}
