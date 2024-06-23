package service

import (
	"context"
	"fmt"
	"github.com/ntc-goer/microservice-examples/kitchen/config"
	kitchenpb "github.com/ntc-goer/microservice-examples/kitchen/proto"
	"time"
)

type KitchenService struct {
	kitchenpb.UnimplementedKitchenServiceServer
	Config *config.Config
}

func NewKitchenService(cfg *config.Config) (*KitchenService, error) {
	return &KitchenService{
		Config: cfg,
	}, nil
}

func (s *KitchenService) VerifyOrder(ctx context.Context, req *kitchenpb.VerifyOrderRequest) (*kitchenpb.VerifyOrderResponse, error) {
	fmt.Println("Verifying order of", req.StoreId)
	time.Sleep(5 * time.Second)
	return &kitchenpb.VerifyOrderResponse{
		IsOk: true,
	}, nil
}
