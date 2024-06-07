package main

import (
	"context"
	pb "github.com/ntc-goer/microservice-examples/proto"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	dp, err := InitializeDependency("consul")
	if err != nil {
		log.Fatalf("fail to init dependency %v", err)
	}
	ctx := context.Background()

	// Setup grpc server
	lis, err := net.Listen("tcp", ":"+dp.Config.GRPCPort)
	if err != nil {
		log.Fatalf("error listening port %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, dp.ServiceImpl)
	pb.RegisterHealthServer(grpcServer, dp.ServiceImpl)
	// Register to discovery service
	instanceId := servicediscovery.GenerateInstanceId(dp.Config.ServiceName)
	if err := dp.ServiceDiscovery.RegisterService(instanceId, dp.Config.ServiceName, dp.Config.GRPCHost, dp.Config.GRPCPort); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer dp.ServiceDiscovery.Deregister(ctx, instanceId)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
