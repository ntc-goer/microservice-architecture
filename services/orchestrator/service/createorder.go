package service

import (
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
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

func (s *CreateOrderService) Run(msg string) {
	return
}
