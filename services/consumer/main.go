package main

import (
	"context"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	dp, err := InitializeDependency("consul")
	if err != nil {
		log.Fatal(err)
	}

	// Setup grpc server
	lis, err := net.Listen("tcp", ":"+dp.Config.GRPCPort)
	if err != nil {
		log.Fatalf("Listen port fail %v", err)
	}
	grpcServer := grpc.NewServer()
	consumerpb.RegisterConsumerServiceServer(grpcServer, dp.ServiceImpl)
	grpc_health_v1.RegisterHealthServer(grpcServer, dp.ServiceImpl)

	// Register to discovery service
	srvDiscovery, err := serviceregistration.GetDiscovery("consul")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	instanceId := serviceregistration.GenerateInstanceId(dp.Config.ConsumerServiceName)
	defer srvDiscovery.Deregister(ctx, instanceId)
	go func(srvd common.DiscoveryI) {
		if err := srvd.RegisterService(instanceId, dp.Config.ConsumerServiceName, serviceregistration.GetCurrentIP(), dp.Config.GRPCPort, common.GRPC_CHECK_TYPE); err != nil {
			log.Fatalf("RegisterService fail: %v", err)
		}
	}(srvDiscovery)

	// Start listen request
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Start server fail")
	}
}
