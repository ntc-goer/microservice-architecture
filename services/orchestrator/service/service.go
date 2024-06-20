package service

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
	"github.com/ntc-goer/microservice-examples/registry/broker"
	"log"
)

type CoreService struct {
	Broker      *broker.Broker
	Health      *HealthService
	CreateOrder *CreateOrderService
	Config      *config.Config
}

func NewCoreService(broker *broker.Broker, healthService *HealthService, createOrderService *CreateOrderService, cfg *config.Config) *CoreService {
	return &CoreService{
		Broker:      broker,
		Health:      healthService,
		CreateOrder: createOrderService,
		Config:      cfg,
	}
}

func (vs *CoreService) StartSubscribe() {
	go func() {
		if err := vs.Broker.QueueSubscribe(vs.CreateOrder.Subject, vs.CreateOrder.Queue, vs.CreateOrder.Run); err != nil {
			log.Printf("Subscribe FAIL %s", err)
		}
		// Keep the connection alive
		select {}
	}()
}

var WireSet = wire.NewSet(NewHealthService, NewCreateOrderService, NewCoreService)
