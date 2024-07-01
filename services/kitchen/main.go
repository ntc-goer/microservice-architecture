package main

import (
	"context"
	kitchenpb "github.com/ntc-goer/microservice-examples/kitchen/proto"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
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

	// Init tracing
	tp, err := initTraceProvider(ctx, dp.Config.Service.KitchenServiceName, "0.1.1", "http://localhost:14268/api/traces")
	if err != nil {
		log.Fatalf("initTraceProvider fail: %v", err)
	}
	defer tp.Shutdown(ctx)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Migrate database
	err = dp.Repository.MigrateDatabase()
	if err != nil {
		log.Fatalf("fail to init dependency %v", err)
	}

	// Setup grpc server
	lis, err := net.Listen("tcp", ":"+dp.Config.ServicePort)
	if err != nil {
		log.Fatalf("Listen port fail %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()))
	kitchenpb.RegisterKitchenServiceServer(grpcServer, dp.Service.Kitchen)
	grpc_health_v1.RegisterHealthServer(grpcServer, dp.Service.Health)

	// Register to discovery service
	srvDiscovery, err := serviceregistration.GetDiscovery("consul")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
	instanceId := serviceregistration.GenerateInstanceId(dp.Config.Service.KitchenServiceName)
	defer srvDiscovery.Deregister(ctx, instanceId)
	go func(srvd common.DiscoveryI) {
		if err := srvd.RegisterService(instanceId, dp.Config.Service.KitchenServiceName, serviceregistration.GetCurrentIP(), dp.Config.ServicePort, common.GRPC_CHECK_TYPE); err != nil {
			log.Fatalf("RegisterService fail: %v", err)
		}
	}(srvDiscovery)

	// Start listen request
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Start server fail")
	}
}
