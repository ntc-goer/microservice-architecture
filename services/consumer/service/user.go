package service

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/consumerservice/config"
	consumerpb "github.com/ntc-goer/microservice-examples/consumerservice/proto"
	"time"
)

type UserService struct {
	consumerpb.UnimplementedConsumerServiceServer
	Config *config.Config
}

func NewUserService(cfg *config.Config) (*UserService, error) {
	return &UserService{
		Config: cfg,
	}, nil
}

func (s *UserService) VerifyUser(ctx context.Context, req *consumerpb.VerifyUserRequest) (*consumerpb.VerifyUserResponse, error) {
	fmt.Println("Verifying user ", req.Id)
	time.Sleep(5 * time.Second)
	return &consumerpb.VerifyUserResponse{
		IsOk: true,
	}, nil
}
