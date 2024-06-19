package main

import (
	"context"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func main() {
	dp, err := InitializeDependency("consul")
	if err != nil {
		log.Fatalf("fail to init dependency %v", err)
	}
	ctx := context.Background()

	// Migrate database
	err = dp.Repository.MigrateDatabase()
	if err != nil {
		log.Fatalf("fail to init dependency %v", err)
	}

	// Setup grpc server
	lis, err := net.Listen("tcp", ":"+dp.Config.ServicePort)
	if err != nil {
		log.Fatalf("error listening port %v", err)
	}
	grpcServer := grpc.NewServer()
	orderpb.RegisterOrderServiceServer(grpcServer, dp.CoreService.ServiceImpl)
	grpc_health_v1.RegisterHealthServer(grpcServer, dp.CoreService.Health)
	// Register to discovery service
	instanceId := serviceregistration.GenerateInstanceId(dp.Config.Service.OrderServiceName)
	if err := dp.ServiceDiscovery.RegisterService(instanceId, dp.Config.Service.OrderServiceName, serviceregistration.GetCurrentIP(), dp.Config.ServicePort, common.GRPC_CHECK_TYPE); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer dp.ServiceDiscovery.Deregister(ctx, instanceId)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
