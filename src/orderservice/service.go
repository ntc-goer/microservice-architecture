package main

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/orderservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type ServiceImpl struct {
	proto.UnimplementedOrderServiceServer
	consumerService proto.ConsumerServiceClient
}

func (s *ServiceImpl) Order(ctx context.Context, orderReq *proto.OrderRequest) (*proto.OrderResponse, error) {
	fmt.Println("Verify User Start")
	result, err := s.consumerService.VerifyUser(ctx, &proto.VerifyUserRequest{Id: orderReq.UserId})
	if err != nil {
		log.Printf("Error when calling the consumer service %v", err)
		return nil, err
	}
	log.Printf("Verify data done with result %v", result.IsOk)
	return &proto.OrderResponse{
		IsOk: result.IsOk,
	}, nil
}

func NewServiceImpl() (*ServiceImpl, error) {
	// Setup connection to other service
	consumerConn, err := grpc.NewClient(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	consumerClient := proto.NewConsumerServiceClient(consumerConn)
	return &ServiceImpl{
		consumerService: consumerClient,
	}, nil
}
