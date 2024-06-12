package service

import (
	"context"
	"fmt"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
)

type ServiceImpl struct {
	orderpb.UnimplementedOrderServiceServer
	SrvDis common.DiscoveryI
	Repo   *repository.Repository
	Config *config.Config
}

func NewServiceImpl(srvDis common.DiscoveryI, repo *repository.Repository, cfg *config.Config) (*ServiceImpl, error) {
	return &ServiceImpl{
		SrvDis: srvDis,
		Repo:   repo,
		Config: cfg,
	}, nil
}

func (s *ServiceImpl) Order(ctx context.Context, orderReq *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	conn, err := grpc.NewClient(fmt.Sprintf("%s/%s", s.Config.LBServiceHost, s.Config.ConsumerServiceName), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := consumerpb.NewConsumerServiceClient(conn)
	result, err := client.VerifyUser(ctx, &consumerpb.VerifyUserRequest{Id: orderReq.UserId})
	if err != nil {
		log.Printf("Error when calling the consumer service %v", err)
		return nil, err
	}
	log.Printf("Verify data done with result %v", result.IsOk)
	return &orderpb.OrderResponse{
		IsOk: result.IsOk,
	}, nil
}

func (s *ServiceImpl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *ServiceImpl) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
