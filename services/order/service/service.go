package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/pkg"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/registry/queue"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/ntc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
)

type Impl struct {
	orderpb.UnimplementedOrderServiceServer
	SrvDis      common.DiscoveryI
	Repo        *repository.Repository
	Config      *config.Config
	LoadBalance *pkg.LB
	Queue       *queue.MsgQueue
}

func NewServiceImpl(srvDis common.DiscoveryI, repo *repository.Repository, cfg *config.Config, lb *pkg.LB, q *queue.MsgQueue) (*Impl, error) {
	return &Impl{
		SrvDis:      srvDis,
		Repo:        repo,
		Config:      cfg,
		LoadBalance: lb,
		Queue:       q,
	}, nil
}

func (s *Impl) Order(ctx context.Context, orderReq *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	//var result *consumerpb.VerifyUserResponse
	//hystrix.ConfigureCommand("order", hystrix.CommandConfig{
	//	// The number of requests will be calculated to determine whether the circuit should be opened or not.
	//	RequestVolumeThreshold: 10,
	//	// The millisecond to determine the request will be timeout or not
	//	Timeout: 1000,
	//	// The millisecond number represent the time the circuit breaker will open until the next test.
	//	SleepWindow: 60000,
	//	// Using RequestVolumeThreshold , calculate the number of request error , if > ErrorPercentThreshold -> The circuit will open
	//	ErrorPercentThreshold: 30,
	//})
	//err := hystrix.Do("CREATE_ORDER", func() error {
	//	conn, err := s.LoadBalance.GetConnection(s.Config.ConsumerServiceName)
	//	if err != nil {
	//		return err
	//	}
	//	client := consumerpb.NewConsumerServiceClient(conn)
	//	result, err = client.VerifyUser(ctx, &consumerpb.VerifyUserRequest{Id: orderReq.UserId})
	//	if err != nil {
	//		log.Printf("Error when calling the consumer service %v", err)
	//		return err
	//	}
	//	return nil
	//}, func(err error) error {
	//	log.Printf("hystrix fallback %s", err.Error())
	//	return err
	//})
	//
	//if err != nil {
	//	return nil, err
	//}
	requestId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	orderId, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	tx, err := s.Repo.Client.Tx(ctx)
	if err != nil {
		return nil, fmt.Errorf("starting a transaction: %w", err)
	}
	// Create an order with APPROVAL_PENDING state
	err = s.Repo.Order.CreatePendingOrder(ctx, orderId, requestId, orderReq.Location, orderReq.UserId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// Create an order with APPROVAL_PENDING state
	dishes := ntc.Map(orderReq.Orders, func(it *orderpb.OrderItem) *repository.DishItem {
		return &repository.DishItem{
			DishId:   "dishId",
			DishName: it.Dish,
			Quantity: int(it.Total),
		}
	})
	err = s.Repo.Dish.CreateDishes(ctx, orderId, dishes)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	// Publish OrderCreate Message to orchestrator
	if err := s.Queue.Connect(s.Config.QueueAddress); err != nil {
		tx.Rollback()
		return nil, err
	}
	defer s.Queue.Close()
	msgBytes, err := json.Marshal(map[string]string{
		"request_id": requestId.String(),
		"order_id":   orderId.String(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = s.Queue.Publish(s.Config.QueueCreateOrderSubject, string(msgBytes))
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &orderpb.OrderResponse{
		RequestId: requestId.String(),
		OrderId:   orderId.String(),
	}, nil
}

func (s *Impl) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (s *Impl) Watch(in *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "health check via Watch not implemented")
}
