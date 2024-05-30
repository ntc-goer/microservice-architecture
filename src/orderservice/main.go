package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ntc-goer/microservice-examples/orderservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

const _SERVICENAME = "orderservice"

func main() {
	// Setup http server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	err := proto.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, ":50000", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}
	go func() {
		log.Printf("starting HTTP/JSON gateway on :8080")
		if err := http.ListenAndServe(":8080", mux); err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	// Setup grpc server
	svImpl, err := NewServiceImpl()
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	lis, err := net.Listen("tcp", ":50000")
	if err != nil {
		log.Fatalf("error listening port %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterOrderServiceServer(grpcServer, svImpl)
	fmt.Println("Order service is running on port :50000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
