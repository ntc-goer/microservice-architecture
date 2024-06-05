package service

import (
	"context"
	"fmt"
	pb "github.com/ntc-goer/microservice-examples/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
)

type ServiceImpl struct {
	pb.UnimplementedOrderServiceServer
	pb.UnimplementedHealthServer
	consumerService pb.ConsumerServiceClient
}

func NewServiceImpl() (*ServiceImpl, error) {
	// Setup connection to other service
	consumerConn, err := grpc.NewClient(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	consumerClient := pb.NewConsumerServiceClient(consumerConn)
	return &ServiceImpl{
		consumerService: consumerClient,
	}, nil
}

func (s *ServiceImpl) Order(ctx context.Context, orderReq *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Println("Verify User Start")
	result, err := s.consumerService.VerifyUser(ctx, &pb.VerifyUserRequest{Id: orderReq.UserId})
	if err != nil {
		log.Printf("Error when calling the consumer service %v", err)
		return nil, err
	}
	log.Printf("Verify data done with result %v", result.IsOk)
	return &pb.OrderResponse{
		IsOk: result.IsOk,
	}, nil
}

func (s *ServiceImpl) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_SERVING}, nil
}

func (s *ServiceImpl) Watch(req *pb.HealthCheckRequest, server pb.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
