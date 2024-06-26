package service

import (
	"context"
	"encoding/json"
	"errors"
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

type CreateOrderStore struct {
	TicketId string
}

func (s *CreateOrderService) Run(msg string) {
	var req OrderMsg
	if err := json.Unmarshal([]byte(msg), &req); err != nil {
		log.Fatalf("Invalid Data Format %s", err)
	}
	wl, err := sagaorchestration.NewWorkflow[*CreateOrderStore]("CREATE_ORDER",
		sagaorchestration.WorkflowConfig[*CreateOrderStore]{
			RequestID: req.RequestId,
			TrackingDB: &sagaorchestration.DB{
				DriverName: "postgres",
				Address:    s.Config.Database.ServerHost,
				Port:       s.Config.Database.ServerPort,
				DBName:     s.Config.Database.DBName,
				UserName:   s.Config.Database.UserName,
				Password:   s.Config.Database.Password,
			},
		})
	if err != nil {
		return
	}
	steps := s.getSagaStep(&req)
	err = wl.RegisterSteps(steps).Start()
	if err != nil {
		// TODO Handle workflow fail
	}
	return
}

func (s *CreateOrderService) getSagaStep(req *OrderMsg) []sagaorchestration.Step[*CreateOrderStore] {
	steps := []sagaorchestration.Step[*CreateOrderStore]{
		// Consumer Service — Verify user can order. \
		{
			Name: "VERIFY_USER",
			ProcessF: func(store *CreateOrderStore) error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
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

				// Update Order TO FAILED
				orderC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.OrderServiceName, orderpb.NewOrderServiceClient)
				// Create a Ticket as CREATE_PENDING state.
				updateOrderStatusRes, err := orderC.UpdateOrderStatusFailed(ctx, &orderpb.UpdateOrderStatusFailedRequest{
					OrderId:   req.OrderId,
					RequestId: req.RequestId,
				})
				if err != nil {
					log.Printf("Error when calling the service %s : %v", s.Config.Service.OrderServiceName, err)
					return err
				}
				if !updateOrderStatusRes.IsOk {
					log.Printf("UpdateOrderStatusFailed %s Fail", req.OrderId)
					return errors.New("UpdateOrderStatusFailed Fail")
				}
				return nil
			},
		},
		// Kitchen Service — Verify order
		// Create a Ticket as CREATE_PENDING state.
		{
			Name: "VERIFY_ORDER",
			ProcessF: func(store *CreateOrderStore) error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
				defer cancel()

				// Get Order Dishes
				orderC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.KitchenServiceName, orderpb.NewDishServiceClient)
				dishRes, err := orderC.GetOrderDish(ctx, &orderpb.GetOrderDishRequest{
					OrderId: req.OrderId,
				})
				if err != nil {
					log.Printf("Error when calling the consumer service %v", err)
					return err
				}
				if len(dishRes.Dishes) == 0 {
					log.Printf("Not found dishes")
					return errors.New("not found dishes")
				}
				// Verify order
				kitchenC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.KitchenServiceName, kitchenpb.NewKitchenServiceClient)
				verifyOrderRes, err := kitchenC.VerifyOrder(ctx, &kitchenpb.VerifyOrderRequest{
					StoreId: req.OrderId,
					Dishes: ntc.Map(dishRes.Dishes, func(d *orderpb.DishItem) *kitchenpb.DishItem {
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
				if !verifyOrderRes.IsOk {
					log.Printf("VerifyOrder %s Fail", req.OrderId)
					return errors.New("VerifyOrder Fail")
				}

				// Create a Ticket as CREATE_PENDING state.
				createPendingTicketResult, err := kitchenC.CreatePendingTicket(ctx, &kitchenpb.CreatePendingTicketRequest{
					RequestId: req.RequestId,
					OrderId:   req.OrderId,
				})
				if err != nil {
					log.Printf("Error when calling the service %s : %v", s.Config.Service.KitchenServiceName, err)
					return err
				}
				if !createPendingTicketResult.IsOk {
					log.Printf("CreatePendingTicket %s Fail", store.TicketId)
					return errors.New("CreatePendingTicket Fail")
				}
				return nil
			},
			CompensatingF: func(store *CreateOrderStore) error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
				defer cancel()
				// Update Ticket TO CANCELED
				if store.TicketId != "" {
					kitchenC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.KitchenServiceName, kitchenpb.NewKitchenServiceClient)
					// Create a Ticket as CREATE_PENDING state.
					cancelTicketResult, err := kitchenC.CancelTicket(ctx, &kitchenpb.CancelTicketRequest{
						TicketId: store.TicketId,
					})
					if err != nil {
						log.Printf("Error when calling the service %s : %v", s.Config.Service.KitchenServiceName, err)
						return err
					}
					if !cancelTicketResult.IsOk {
						log.Printf("CancelTicket %s Fail", store.TicketId)
						return errors.New("CancelTicket Fail")
					}
					return nil
				}
				return nil
			},
		},
		// Accounting Service — Verify user's credit card.
		{
			Name: "VERIFY_USER_CREDIT_CARD",
			ProcessF: func(store *CreateOrderStore) error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
				defer cancel()

				// Verify credit card
				client, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.AccountingServiceName, accountingpb.NewAccountingServiceClient)
				result, err := client.VerifyCreditCard(ctx, &accountingpb.VerifyCreditCardRequest{
					UserId: req.UserId,
				})
				if err != nil {
					log.Printf("Error when calling the service %s : %v", s.Config.Service.AccountingServiceName, err)
					return err
				}
				if !result.IsOk {
					log.Printf("VerifyCreditCard of  %s Fail", req.UserId)
					return errors.New("VerifyCreditCard Fail")
				}

				return nil
			},
		},
		// Kitchen Service — Change ticket's state to AWAITING_ACCEPTANCE.\
		{
			Name: "UPDATE_TICKET_STATE_TO_AWAITING_ACCEPTANCE",
			ProcessF: func(store *CreateOrderStore) error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
				defer cancel()
				// Update Ticket TO CANCELED
				if store.TicketId != "" {
					kitchenC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.KitchenServiceName, kitchenpb.NewKitchenServiceClient)
					// Create a Ticket as CREATE_PENDING state.
					acceptTicketResult, err := kitchenC.AcceptTicket(ctx, &kitchenpb.AcceptTicketRequest{
						TicketId: store.TicketId,
					})
					if err != nil {
						log.Printf("Error when calling the service %s : %v", s.Config.Service.KitchenServiceName, err)
						return err
					}
					if !acceptTicketResult.IsOk {
						log.Printf("AcceptTicket %s Fail", store.TicketId)
						return errors.New("AcceptTicket Fail")
					}
					return nil
				}
				return nil
			},
		},
		// Order Service — Change order state to APPROVED.
		{
			Name: "UPDATE_ORDER_STATE_TO_APPROVED",
			ProcessF: func(store *CreateOrderStore) error {
				ctx := context.Background()
				ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
				defer cancel()

				// Update order_status APPROVED
				orderC, err := pkg.GetGRPCClient(s.Config.Service.LBServiceHost, s.Config.Service.OrderServiceName, orderpb.NewOrderServiceClient)
				approveOrderRes, err := orderC.ApproveOrder(ctx, &orderpb.ApproveOrderRequest{
					RequestId: req.RequestId,
					OrderId:   req.OrderId,
				})
				if err != nil {
					log.Printf("Error when calling the consumer service %v", err)
					return err
				}
				if !approveOrderRes.IsOk {
					log.Printf("ApproveOrder %s Fail", req.OrderId)
					return errors.New("ApproveOrder Fail")
				}
				return nil
			},
		},
	}
	return steps
}
