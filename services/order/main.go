package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/service"
	pb "github.com/ntc-goer/microservice-examples/proto"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	// Setup http server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err = pb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", cfg.Host.GRPCPort), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}
	go func() {
		log.Printf("starting HTTP/JSON gateway on :8080")
		if err := http.ListenAndServe(":"+cfg.Host.HTTPPort, mux); err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	// Setup grpc server
	svImpl, err := service.NewServiceImpl()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	lis, err := net.Listen("tcp", ":"+cfg.Host.GRPCPort)
	if err != nil {
		log.Fatalf("error listening port %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, svImpl)
	// Register to discovery service
	srvDiscovery, err := servicediscovery.NewServiceDiscovery()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	srvDiscovery.RegisterService(cfg.Host.ServiceId, cfg.Host.ServiceName, cfg.Host.GRPCHost, cfg.Host.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
