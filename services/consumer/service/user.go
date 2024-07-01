package service

import (
	"context"
	"github.com/ntc-goer/microservice-examples/consumerservice/config"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"time"
)

type UserService struct {
	consumerpb.UnimplementedConsumerServiceServer
	Config *config.Config
	Trace  trace.Tracer
}

func NewUserService(cfg *config.Config) (*UserService, error) {
	return &UserService{
		Config: cfg,
		Trace:  otel.Tracer("UserService"),
	}, nil
}

func (s *UserService) VerifyUser(ctx context.Context, req *consumerpb.VerifyUserRequest) (*consumerpb.VerifyUserResponse, error) {
	ctx, span := s.Trace.Start(ctx, "UserService.VerifyUser")
	defer span.End()
	// Add attributes to the span
	span.SetAttributes(attribute.String("UserId", req.Id))

	time.Sleep(5 * time.Second)
	return &consumerpb.VerifyUserResponse{
		IsOk: true,
	}, nil
}
