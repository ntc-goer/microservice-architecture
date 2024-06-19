package service

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/registry/broker"
)

type CoreService struct {
	Broker      *broker.Broker
	Health      *HealthService
	CreateOrder *CreateOrderService
}

func NewCoreService(broker *broker.Broker, healthService *HealthService, createOrderService *CreateOrderService) *CoreService {
	return &CoreService{
		Broker:      broker,
		Health:      healthService,
		CreateOrder: createOrderService,
	}
}

func (vs *CoreService) StartSubscribe() {
	go func() {
		vs.Broker.QueueSubscribe(vs.CreateOrder.Subject, vs.CreateOrder.Queue, vs.CreateOrder.Run)
	}()
}

var WireSet = wire.NewSet(NewHealthService, NewCreateOrderService, NewCoreService)
