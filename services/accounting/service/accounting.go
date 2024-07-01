package service

import (
	"context"
	"github.com/ntc-goer/microservice-examples/accounting/config"
	accountingpb "github.com/ntc-goer/microservice-examples/accounting/proto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type AccountingService struct {
	accountingpb.UnimplementedAccountingServiceServer
	Config *config.Config
	Trace  trace.Tracer
}

func NewAccountingService(cfg *config.Config) (*AccountingService, error) {
	return &AccountingService{
		Config: cfg,
		Trace:  otel.Tracer("AccountingService"),
	}, nil
}

func (s *AccountingService) VerifyCreditCard(ctx context.Context, req *accountingpb.VerifyCreditCardRequest) (*accountingpb.VerifyCreditCardResponse, error) {
	ctx, span := s.Trace.Start(ctx, "AccountingService.VerifyCreditCard")
	defer span.End()
	// Add attributes to the span
	span.SetAttributes(attribute.String("UserId", req.UserId))

	time.Sleep(5 * time.Second)
	return &accountingpb.VerifyCreditCardResponse{
		IsOk: true,
	}, nil
}
