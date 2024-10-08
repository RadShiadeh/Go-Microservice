package main

import (
	"context"
	"get-price/proto"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

func CreateGRPCServerAndRun(listenAddr string, svc PriceGetter) error {

	grpcPriceGetter := NewGrpcPriceService(svc)

	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	proto.RegisterPriceGetterServer(server, grpcPriceGetter)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	return server.Serve(lis)
}

type GrpcPriceGetterServer struct {
	svc PriceGetter
	proto.UnimplementedPriceGetterServer
}

func NewGrpcPriceService(svc PriceGetter) *GrpcPriceGetterServer {

	return &GrpcPriceGetterServer{
		svc: svc,
	}
}

func (server *GrpcPriceGetterServer) GetPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {

	reqID := rand.Intn(100000)
	ctx = context.WithValue(ctx, "requestID", reqID)

	price, err := server.svc.GetPrice(ctx, req.Key, req.Currency)
	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Price: float32(price),
		Key:   req.Key,
	}

	return resp, nil
}
