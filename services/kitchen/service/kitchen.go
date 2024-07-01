package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ntc-goer/microservice-examples/kitchen/config"
	kitchenpb "github.com/ntc-goer/microservice-examples/kitchen/proto"
	"github.com/ntc-goer/microservice-examples/kitchen/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type KitchenService struct {
	kitchenpb.UnimplementedKitchenServiceServer
	Config *config.Config
	Repo   *repository.Repository
	Trace  trace.Tracer
}

func NewKitchenService(cfg *config.Config, repo *repository.Repository) (*KitchenService, error) {
	return &KitchenService{
		Config: cfg,
		Repo:   repo,
		Trace:  otel.Tracer("KitchenService"),
	}, nil
}

func (s *KitchenService) VerifyOrder(ctx context.Context, req *kitchenpb.VerifyOrderRequest) (*kitchenpb.VerifyOrderResponse, error) {
	ctx, span := s.Trace.Start(ctx, "KitchenService.VerifyOrder")
	defer span.End()
	// Add attributes to the span
	dishesByte, _ := json.Marshal(req.Dishes)
	span.SetAttributes(
		attribute.String("StoreId", req.StoreId),
		attribute.String("Dishes", string(dishesByte)))

	fmt.Println("Verifying order of", req.StoreId)
	time.Sleep(5 * time.Second)
	return &kitchenpb.VerifyOrderResponse{
		IsOk: true,
	}, nil
}

func (s *KitchenService) CreatePendingTicket(ctx context.Context, req *kitchenpb.CreatePendingTicketRequest) (*kitchenpb.CreatePendingTicketResponse, error) {
	ctx, span := s.Trace.Start(ctx, "KitchenService.CreatePendingTicket")
	defer span.End()
	// Add attributes to the span
	span.SetAttributes(
		attribute.String("OrderId", req.OrderId),
		attribute.String("RequestId", req.RequestId))

	ticket, err := s.Repo.Ticket.CreatePendingTicket(ctx, req.OrderId, req.RequestId)
	if err != nil {
		return nil, err
	}

	return &kitchenpb.CreatePendingTicketResponse{IsOk: true, TicketId: ticket.ID.String()}, nil
}

func (s *KitchenService) AcceptTicket(ctx context.Context, req *kitchenpb.AcceptTicketRequest) (*kitchenpb.AcceptTicketResponse, error) {
	ctx, span := s.Trace.Start(ctx, "KitchenService.AcceptTicket")
	defer span.End()
	// Add attributes to the span
	span.SetAttributes(
		attribute.String("TicketId", req.TicketId))

	_, err := s.Repo.Ticket.AcceptTicket(ctx, req.TicketId)
	if err != nil {
		return nil, err
	}

	return &kitchenpb.AcceptTicketResponse{IsOk: true}, nil
}

func (s *KitchenService) CancelTicket(ctx context.Context, req *kitchenpb.CancelTicketRequest) (*kitchenpb.CancelTicketResponse, error) {
	ctx, span := s.Trace.Start(ctx, "KitchenService.CancelTicket")
	defer span.End()
	// Add attributes to the span
	span.SetAttributes(
		attribute.String("TicketId", req.TicketId))

	_, err := s.Repo.Ticket.CancelTicket(ctx, req.TicketId)
	if err != nil {
		return nil, err
	}

	return &kitchenpb.CancelTicketResponse{IsOk: true}, nil
}
