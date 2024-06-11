package service

import (
	"context"
	"errors"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type ServiceImpl struct {
	orderpb.UnimplementedOrderServiceServer
	SrvDis common.DiscoveryI
	Repo   *repository.Repository
}

func NewServiceImpl(srvDis common.DiscoveryI, repo *repository.Repository) (*ServiceImpl, error) {
	return &ServiceImpl{
		SrvDis: srvDis,
		Repo:   repo,
	}, nil
}

func (s *ServiceImpl) GetActiveService(ctx context.Context, serviceName string) (*grpc.ClientConn, error) {
	srvs, err := s.SrvDis.Discover(ctx, serviceName)
	if err != nil {
		log.Printf("Discover service fail %v", err)
		return nil, err
	}
	if len(srvs) == 0 {
		err := errors.New("not Found active service")
		log.Printf("Not found active service %v", err)
		return nil, err
	}
	// Setup connection to other service
	consumerConn, err := grpc.NewClient(":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return consumerConn, nil

}
func (s *ServiceImpl) Order(ctx context.Context, orderReq *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	conn, err := s.GetActiveService(ctx, "consumer")
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

func (s *ServiceImpl) Check(ctx context.Context, e *emptypb.Empty) (*orderpb.HealthCheckResponse, error) {
	return &orderpb.HealthCheckResponse{Status: orderpb.HealthCheckResponse_SERVING}, nil
}

func (s *ServiceImpl) Watch(req *orderpb.HealthCheckRequest, server orderpb.OrderService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
