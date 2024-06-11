package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	// Setup http server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	// Register endpoint
	// OrderService
	_ = orderpb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", "50000"), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	_ = consumerpb.RegisterConsumerServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", "50001"), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	log.Printf("starting HTTP/JSON gateway on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
