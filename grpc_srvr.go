package main

import (
	"context"
	"get-price/proto"
	"net"

	"google.golang.org/grpc"
)

func CreateGRPCServerAndRun(listenAddr string, svc PriceGetter) error {

	grpcPriceGetter := NewGrpcPriceService(svc)

	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		return err
	}

	options := []grpc.ServerOption{}
	server := grpc.NewServer(options...)

	proto.RegisterPriceGetterServer(server, grpcPriceGetter)

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

	price, err := server.svc.GetPrice(ctx, req.Key, req.Currency)
	if err != nil {
		return nil, err
	}

	resp := &proto.PriceResponse{
		Price: price,
		Key:   req.Key,
	}

	return resp, nil
}
