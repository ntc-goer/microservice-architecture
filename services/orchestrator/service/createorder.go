package service

import (
	"encoding/json"
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
	"log"
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

type OrderMsg struct {
	UserId    string `json:"user_id"`
	OrderId   string `json:"order_id"`
	RequestId string `json:"request_id"`
}

func (s *CreateOrderService) Run(msg string) {
	var req OrderMsg
	if err := json.Unmarshal([]byte(msg), &req); err != nil {
		log.Fatalf("Invalid Data Format %s", err)
	}
	// TODO Define SAGA
	// Consumer Service — Verify an user can order. \
	// Kitchen Service — Verify order và create a Ticket as CREATE_PENDING state.\
	// Accounting Service — Verify user's credit card.\
	// Kitchen Service — Change ticket's state to AWAITING_ACCEPTANCE.\
	// Order Service — Change order state to APPROVED.
	return
}
