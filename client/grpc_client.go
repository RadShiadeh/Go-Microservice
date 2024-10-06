package client

import (
	"fmt"
	"get-price/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClient(addr string) (proto.PriceGetterClient, error) {

	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc client error: %v", err)
	}

	client := proto.NewPriceGetterClient(conn)

	return client, nil
}
