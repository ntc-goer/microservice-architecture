package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/registry/broker"
	"github.com/ntc-goer/ntc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type OrderService struct {
	orderpb.UnimplementedOrderServiceServer
	Repo   *repository.Repository
	Queue  *broker.Broker
	Config *config.Config
	Trace  trace.Tracer
}

func NewOrderService(repo *repository.Repository, cfg *config.Config, q *broker.Broker) *OrderService {
	return &OrderService{
		Repo:   repo,
		Config: cfg,
		Queue:  q,
		Trace:  otel.Tracer("OrderService"),
	}
}

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

func (s *OrderService) Order(ctx context.Context, orderReq *orderpb.OrderRequest) (*orderpb.OrderResponse, error) {
	ctx, span := s.Trace.Start(ctx, "OrderService.StartOrder")
	defer span.End()
	traceId := span.SpanContext().TraceID()

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
	ord, err := s.Repo.Order.CreatePendingOrder(ctx, orderId, requestId, orderReq.Location, orderReq.UserId)
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
	if err := s.Queue.Connect(s.Config.Broker.Address); err != nil {
		tx.Rollback()
		return nil, err
	}
	defer s.Queue.Close()
	msgBytes, err := json.Marshal(map[string]string{
		"user_id":    ord.UserID,
		"request_id": ord.RequestID.String(),
		"order_id":   ord.ID.String(),
	})
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = s.Queue.Publish(ctx, s.Config.Broker.Subject.CreateOrder, msgBytes)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &orderpb.OrderResponse{
		RequestId: ord.RequestID.String(),
		OrderId:   ord.ID.String(),
		Status:    string(ord.Status),
		TraceId:   traceId.String(),
	}, nil
}

func (s *OrderService) ApproveOrder(ctx context.Context, req *orderpb.ApproveOrderRequest) (*orderpb.ApproveOrderResponse, error) {
	ctx, span := s.Trace.Start(ctx, "OrderService.ApproveOrder")
	defer span.End()

	orderUUID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}
	requestUUID, err := uuid.Parse(req.RequestId)
	if err != nil {
		return nil, err
	}
	_, err = s.Repo.Order.ApproveOrder(ctx, orderUUID, requestUUID)
	if err != nil {
		return nil, err
	}
	return &orderpb.ApproveOrderResponse{
		IsOk: true,
	}, nil
}

func (s *OrderService) UpdateOrderStatusFailed(ctx context.Context, req *orderpb.UpdateOrderStatusFailedRequest) (*orderpb.UpdateOrderStatusFailedResponse, error) {
	ctx, span := s.Trace.Start(ctx, "OrderService.UpdateOrderStatusFailed")
	defer span.End()

	orderUUID, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}
	requestUUID, err := uuid.Parse(req.RequestId)
	if err != nil {
		return nil, err
	}
	_, err = s.Repo.Order.UpdateOrderStatusFailed(ctx, orderUUID, requestUUID)
	if err != nil {
		return nil, err
	}
	return &orderpb.UpdateOrderStatusFailedResponse{
		IsOk: true,
	}, nil
}
