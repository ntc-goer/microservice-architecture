package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"github.com/ntc-goer/microservice-examples/gateway/config"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
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
	mux.HandlePath("GET", "/health", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Write([]byte("OK"))
	})

	// Register to discovery service
	serviceDiscovery, _ := serviceregistration.GetDiscovery("consul")
	instanceId := serviceregistration.GenerateInstanceId(cfg.Service.GatewayServiceName)
	if err := serviceDiscovery.RegisterService(instanceId, cfg.Service.GatewayServiceName, serviceregistration.GetCurrentIP(), cfg.ServicePort, common.HTTP_CHECK_TYPE); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer serviceDiscovery.Deregister(ctx, instanceId)

	// Register endpoint
	// OrderService
	err = orderpb.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s", cfg.Service.LBServiceHost), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		return
	}
	// ConsumerService
	_ = consumerpb.RegisterConsumerServiceHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s", cfg.Service.LBServiceHost), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		return
	}

	log.Printf("starting HTTP/JSON gateway on " + cfg.ServicePort)
	if err := http.ListenAndServe(":"+cfg.ServicePort, mux); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}
