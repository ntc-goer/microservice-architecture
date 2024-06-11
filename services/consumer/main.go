package main

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/consumerservice/config"
	pb "github.com/ntc-goer/microservice-examples/proto"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type ServerImpl struct {
	pb.UnimplementedConsumerServiceServer
}

func (s *ServerImpl) VerifyUser(ctx context.Context, req *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error) {
	fmt.Printf("Verifing user with ID %s", req.Id)
	time.Sleep(5 * time.Second)
	fmt.Printf("Verify user with ID %s valid", req.Id)
	return &pb.VerifyUserResponse{
		IsOk: true,
	}, nil
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
	pb.RegisterConsumerServiceServer(grpcServer, &ServerImpl{})

	// Register to discovery service
	instanceId := serviceregistration.GenerateInstanceId(cfg.ServiceName)
	srvDiscovery, err := consul.NewRegistry()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if err := srvDiscovery.RegisterService(instanceId, cfg.ServiceName, cfg.GRPCHost, cfg.GRPCPort, "http://host.docker.internal:8080/consumer/health"); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ticker.C:
				srvDiscovery.HealthCheck(instanceId)
			}
		}
	}()
	defer srvDiscovery.Deregister(ctx, instanceId)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Start server fail")
	}
}
