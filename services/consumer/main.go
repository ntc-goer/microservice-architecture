package main

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/consumerservice/config"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"time"
)

type ServiceImpl struct {
	consumerpb.UnimplementedConsumerServiceServer
}

func (s *ServiceImpl) VerifyUser(ctx context.Context, req *consumerpb.VerifyUserRequest) (*consumerpb.VerifyUserResponse, error) {
	fmt.Printf("Verifing user with ID %s", req.Id)
	time.Sleep(5 * time.Second)
	fmt.Printf("Verify user with ID %s valid", req.Id)
	return &consumerpb.VerifyUserResponse{
		IsOk: true,
	}, nil
}

func (s *ServiceImpl) Check(ctx context.Context, e *emptypb.Empty) (*consumerpb.HealthCheckResponse, error) {
	return &consumerpb.HealthCheckResponse{Status: consumerpb.HealthCheckResponse_SERVING}, nil
}

func (s *ServiceImpl) Watch(req *consumerpb.HealthCheckRequest, server consumerpb.ConsumerService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}

func main() {
	// Setup http server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Load fail")
	}
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Listen port fail %v", err)
	}
	grpcServer := grpc.NewServer()
	consumerpb.RegisterConsumerServiceServer(grpcServer, &ServiceImpl{})

	// Register to discovery service
	instanceId := serviceregistration.GenerateInstanceId(cfg.ServiceName)
	srvDiscovery, err := consul.NewRegistry()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if err := srvDiscovery.RegisterService(instanceId, cfg.ServiceName, serviceregistration.GetCurrentIP(), cfg.GRPCPort, "http://host.docker.internal:8080/consumer/health"); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer srvDiscovery.Deregister(ctx, instanceId)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Start server fail")
	}
}
