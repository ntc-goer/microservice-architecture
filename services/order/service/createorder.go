package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/ntc"
)

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
	if err := s.Queue.Connect(s.Config.Broker.Address); err != nil {
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
	err = s.Queue.Publish(s.Config.Broker.Subject.CreateOrder, string(msgBytes))
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
