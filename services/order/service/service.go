package service

import (
	"context"
	"github.com/afex/hystrix-go/hystrix"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/pkg"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type Impl struct {
	orderpb.UnimplementedOrderServiceServer
	SrvDis      common.DiscoveryI
	Repo        *repository.Repository
	Config      *config.Config
	LoadBalance *pkg.LB
}

func NewServiceImpl(srvDis common.DiscoveryI, repo *repository.Repository, cfg *config.Config, lb *pkg.LB) (*Impl, error) {
	return &Impl{
		SrvDis:      srvDis,
		Repo:        repo,
		Config:      cfg,
		LoadBalance: lb,
	}, nil
}

func (s *Impl) Order(ctx context.Context, orderReq *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	var result *consumerpb.VerifyUserResponse
	hystrix.ConfigureCommand("order", hystrix.CommandConfig{
		// The number of requests will be calculated to determine whether the circuit should be opened or not.
		RequestVolumeThreshold: 10,
		// The millisecond to determine the request will be timeout or not
		Timeout: 1000,
		// The millisecond number represent the time the circuit breaker will open until the next test.
		SleepWindow: 60000,
		// Using RequestVolumeThreshold , calculate the number of request error , if > ErrorPercentThreshold -> The circuit will open
		ErrorPercentThreshold: 30,
	})
	err := hystrix.Do("order", func() error {
		log.Printf("Start Order")
		time.Sleep(3 * time.Second)
		//conn, err := s.LoadBalance.GetConnection(s.Config.ConsumerServiceName)
		//if err != nil {
		//	return err
		//}
		//client := consumerpb.NewConsumerServiceClient(conn)
		//result, err = client.VerifyUser(ctx, &consumerpb.VerifyUserRequest{Id: orderReq.UserId})
		//if err != nil {
		//	log.Printf("Error when calling the consumer service %v", err)
		//	return err
		//}
		return nil
	}, func(err error) error {
		log.Printf("hystrix fallback %s", err.Error())
		return err
	})

	if err != nil {
		return nil, err
	}

	log.Printf("Verify data done with result %v", result.IsOk)
	return &orderpb.OrderResponse{
		IsOk: result.IsOk,
	}, nil
}

func (s *Impl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *Impl) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
