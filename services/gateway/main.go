package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"

	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"github.com/ntc-goer/microservice-examples/gateway/config"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		return
	}
	// Setup http server
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	mux.HandlePath("GET", "/gateway/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("OK"))
	})
	// Register endpoint
	// OrderService
	_ = orderpb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", "50000"), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	// ConsumerService
	_ = consumerpb.RegisterConsumerServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%s", "50001"), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	// Register to discovery service
	serviceDiscovery, _ := serviceregistration.GetDiscovery("consul")
	instanceId := serviceregistration.GenerateInstanceId(cfg.ServiceName)
	if err := serviceDiscovery.RegisterService(instanceId, cfg.ServiceName, serviceregistration.GetCurrentIP(), cfg.HttpPort, "http://host.docker.internal:8080/gateway/health"); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer serviceDiscovery.Deregister(ctx, instanceId)
	log.Printf("starting HTTP/JSON gateway on " + cfg.HttpPort)
	if err := http.ListenAndServe(":"+cfg.HttpPort, mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
