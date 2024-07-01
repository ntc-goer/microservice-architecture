package main

import (
	"context"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
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

	// Init tracing
	tp, err := initTraceProvider(ctx, dp.Config.Service.OrchestratorServiceName, "0.1.1", "http://localhost:14268/api/traces")
	if err != nil {
		log.Fatalf("initTraceProvider fail: %v", err)
	}
	defer tp.Shutdown(ctx)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Connect to Broker
	if err := dp.Broker.Connect(dp.Config.Broker.Address); err != nil {
		log.Fatalf("BrokerConnect fail: %v", err)
	}
	defer dp.Broker.Close()

	// Start listening broker
	dp.CoreService.StartSubscribe(ctx)

	// Setup grpc server
	lis, err := net.Listen("tcp", ":"+dp.Config.ServicePort)
	if err != nil {
		log.Fatalf("error listening port %v", err)
	}
	grpcServer := grpc.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, dp.CoreService.Health)

	// Register to discovery service
	instanceId := serviceregistration.GenerateInstanceId(dp.Config.Service.OrchestratorServiceName)
	if err := dp.ServiceDiscovery.RegisterService(instanceId, dp.Config.Service.OrchestratorServiceName, serviceregistration.GetCurrentIP(), dp.Config.ServicePort, common.GRPC_CHECK_TYPE); err != nil {
		log.Fatalf("RegisterService fail: %v", err)
	}
	defer dp.ServiceDiscovery.Deregister(ctx, instanceId)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to start gRPC server: %v", err)
	}
}
