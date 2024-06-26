package service

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/kitchen/config"
	kitchenpb "github.com/ntc-goer/microservice-examples/kitchen/proto"
	"github.com/ntc-goer/microservice-examples/kitchen/repository"
	"time"
)

type KitchenService struct {
	kitchenpb.UnimplementedKitchenServiceServer
	Config *config.Config
	Repo   *repository.Repository
}

func NewKitchenService(cfg *config.Config, repo *repository.Repository) (*KitchenService, error) {
	return &KitchenService{
		Config: cfg,
		Repo:   repo,
	}, nil
}

func (s *KitchenService) VerifyOrder(ctx context.Context, req *kitchenpb.VerifyOrderRequest) (*kitchenpb.VerifyOrderResponse, error) {
	fmt.Println("Verifying order of", req.StoreId)
	time.Sleep(5 * time.Second)
	return &kitchenpb.VerifyOrderResponse{
		IsOk: true,
	}, nil
}

func (s *KitchenService) CreatePendingTicket(ctx context.Context, req *kitchenpb.CreatePendingTicketRequest) (*kitchenpb.CreatePendingTicketResponse, error) {
	_, err := s.Repo.Ticket.CreatePendingTicket(ctx, req.OrderId, req.RequestId)
	if err != nil {
		return nil, err
	}

	return &kitchenpb.CreatePendingTicketResponse{IsOk: true}, nil
}

func (s *KitchenService) AcceptTicket(ctx context.Context, req *kitchenpb.AcceptTicketRequest) (*kitchenpb.AcceptTicketResponse, error) {
	_, err := s.Repo.Ticket.AcceptTicket(ctx, req.TicketId)
	if err != nil {
		return nil, err
	}

	return &kitchenpb.AcceptTicketResponse{IsOk: true}, nil
}

func (s *KitchenService) CancelTicket(ctx context.Context, req *kitchenpb.CancelTicketRequest) (*kitchenpb.CancelTicketResponse, error) {
	_, err := s.Repo.Ticket.CancelTicket(ctx, req.TicketId)
	if err != nil {
		return nil, err
	}

	return &kitchenpb.CancelTicketResponse{IsOk: true}, nil
}
