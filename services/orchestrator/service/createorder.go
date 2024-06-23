package service

import (
	"context"
	"encoding/json"
	accountingpb "github.com/ntc-goer/microservice-examples/accounting/proto"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	kitchenpb "github.com/ntc-goer/microservice-examples/kitchen/proto"
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
	"github.com/ntc-goer/microservice-examples/orchestrator/pkg"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	sagaorchestration "github.com/ntc-goer/microservice-examples/registry/sagaorchestation"
	"github.com/ntc-goer/ntc"
	"log"
	"time"
)

type CreateOrderService struct {
	Subject string
	Queue   string
	Config  *config.Config
}

func NewCreateOrderService(cfg *config.Config) *CreateOrderService {
	return &CreateOrderService{
		Subject: cfg.Broker.Subject.CreateOrder,
		Queue:   cfg.Broker.Queue.Orchestrator,
		Config:  cfg,
	}
}

type OrderMsg struct {
	UserId    string `json:"user_id"`
	OrderId   string `json:"order_id"`
	RequestId string `json:"request_id"`
}

func (s *CreateOrderService) Run(msg string) {
	var req OrderMsg
	if err := json.Unmarshal([]byte(msg), &req); err != nil {
		log.Fatalf("Invalid Data Format %s", err)
	}
	wl := sagaorchestration.NewWorkflow("CREATE_ORDER")
	steps := s.getSagaStep(&req)
	err := wl.RegisterSteps(steps).Start()
	if err != nil {
		// TODO Handle workflow fail
	}

	return
}

func (s *CreateOrderService) getSagaStep(req *OrderMsg) []sagaorchestration.Step {
	steps := []sagaorchestration.Step{
		// Consumer Service — Verify an user can order. \
		{
			Name: "VERIFY_USER",
			ProcessF: func() error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				client, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.ConsumerServiceName, consumerpb.NewConsumerServiceClient)
				result, err := client.VerifyUser(ctx, &consumerpb.VerifyUserRequest{Id: req.UserId})
				if err != nil {
					log.Printf("Error when calling the consumer service %v", err)
					return err
				}
				if !result.IsOk {
					log.Printf("Invalid user %s", req.UserId)
					return err
				}
				return nil
			},
		},
		// Kitchen Service — Verify order
		// Create a Ticket as CREATE_PENDING state.
		{
			Name: "VERIFY_ORDER",
			ProcessF: func() error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				// Get Order Dishes
				orderC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.KitchenServiceName, orderpb.NewOrderServiceClient)
				dishRes, err := orderC.GetOrderDish(ctx, &orderpb.GetOrderDishRequest{
					OrderId: req.OrderId,
				})
				if err != nil {
					log.Printf("Error when calling the consumer service %v", err)
					return err
				}

				// Verify order
				kitchenC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.KitchenServiceName, kitchenpb.NewKitchenServiceClient)
				result, err := kitchenC.VerifyOrder(ctx, &kitchenpb.VerifyOrderRequest{
					StoreId: req.OrderId,
					Dishes: ntc.Map(dishRes.Dishes, func(d *orderpb.OrderItem) *kitchenpb.DishItem {
						return &kitchenpb.DishItem{
							DishId: d.DishId,
							Dish:   d.Dish,
							Total:  d.Total,
						}
					}),
				})
				if err != nil {
					log.Printf("Error when calling the consumer service %v", err)
					return err
				}
				if !result.IsOk {
					log.Printf("Invalid user %s", req.UserId)
					return err
				}

				// Create a Ticket as CREATE_PENDING state.
				return nil
			},
			CompensatingF: func() error {
				// Update Ticket TO CANCELED
				return nil
			},
		},
		// Accounting Service — Verify user's credit card.
		{
			Name: "VERIFY_USER_CREDIT_CARD",
			ProcessF: func() error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
				defer cancel()

				// Verify credit card
				client, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.AccountingServiceName, accountingpb.NewAccountingServiceClient)
				result, err := client.VerifyCreditCard(ctx, &accountingpb.VerifyCreditCardRequest{
					UserId: req.UserId,
				})
				if err != nil {
					log.Printf("Error when calling the consumer service %v", err)
					return err
				}
				if !result.IsOk {
					log.Printf("Invalid user %s", req.UserId)
					return err
				}

				return nil
			},
		},
		// Kitchen Service — Change ticket's state to AWAITING_ACCEPTANCE.\
		{
			Name: "UPDATE_TICKET_STATE_TO_AWAITING_ACCEPTANCE",
			ProcessF: func() error {
				return nil
			},
			CompensatingF: func() error {
				return nil
			},
		},
		// Order Service — Change order state to APPROVED.
		{
			Name: "UPDATE_ORDER_STATE_TO_APPROVED",
			ProcessF: func() error {
				return nil
			},
			CompensatingF: func() error {
				return nil
			},
		},
	}
	return steps
}
