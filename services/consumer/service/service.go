package service

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/consumerservice/config"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"time"
)

type Impl struct {
	consumerpb.UnimplementedConsumerServiceServer
	SrvDis common.DiscoveryI
	Repo   *repository.Repository
	Config *config.Config
}

func NewServiceImpl(srvDis common.DiscoveryI, repo *repository.Repository, cfg *config.Config) (*Impl, error) {
	return &Impl{
		SrvDis: srvDis,
		Repo:   repo,
		Config: cfg,
	}, nil
}

func (s *Impl) VerifyUser(ctx context.Context, req *consumerpb.VerifyUserRequest) (*consumerpb.VerifyUserResponse, error) {
	fmt.Printf("Verifing user with ID %s", req.Id)
	time.Sleep(5 * time.Second)
	fmt.Printf("Verify user with ID %s valid", req.Id)
	return &consumerpb.VerifyUserResponse{
		IsOk: true,
	}, nil
}

func (s *Impl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *Impl) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
