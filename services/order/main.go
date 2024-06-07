package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	pb "github.com/ntc-goer/microservice-examples/proto"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	dp, err := InitializeDependency("consul")
	if err != nil {
		log.Fatalf("fail to init dependency %v", err)
	}
	// Setup http server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err = pb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", dp.Config.GRPCPort), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}
	pb.RegisterHealthHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", dp.Config.GRPCPort), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}
	go func() {
		log.Printf("starting HTTP/JSON gateway on :8080")
		if err := http.ListenAndServe(":"+dp.Config.HTTPPort, mux); err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

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
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	if err := dp.ServiceDiscovery.RegisterService(instanceId, dp.Config.ServiceName, dp.Config.GRPCHost, dp.Config.GRPCPort); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	//go func() {
	//	ticker := time.NewTicker(3 * time.Second)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			dp.ServiceDiscovery.HealthCheck(instanceId)
	//		}
	//	}
	//}()
	defer dp.ServiceDiscovery.Deregister(ctx, instanceId)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
