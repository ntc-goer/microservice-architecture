package service

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/accounting/config"
	accountingpb "github.com/ntc-goer/microservice-examples/accounting/proto"
	"time"
)

type AccountingService struct {
	accountingpb.UnimplementedAccountingServiceServer
	Config *config.Config
}

func NewAccountingService(cfg *config.Config) (*AccountingService, error) {
	return &AccountingService{
		Config: cfg,
	}, nil
}

func (s *AccountingService) VerifyCreditCard(ctx context.Context, req *accountingpb.VerifyCreditCardRequest) (*accountingpb.VerifyCreditCardResponse, error) {
	fmt.Println("Verifying creadit card of user ", req.UserId)
	time.Sleep(5 * time.Second)
	return &accountingpb.VerifyCreditCardResponse{
		IsOk: true,
	}, nil
}
