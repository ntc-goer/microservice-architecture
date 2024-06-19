package service

import (
	"context"
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
	"github.com/ntc-goer/microservice-examples/orchestrator/pkg"
	"github.com/ntc-goer/microservice-examples/registry/broker"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type Impl struct {
	Config      *config.Config
	LoadBalance *pkg.LB
	Broker      *broker.Broker
}

func NewServiceImpl(cfg *config.Config, lb *pkg.LB, q *broker.Broker) (*Impl, error) {
	return &Impl{
		Config:      cfg,
		LoadBalance: lb,
		Broker:      q,
	}, nil
}

func (s *Impl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *Impl) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
