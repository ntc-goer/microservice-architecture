package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/kitchen/ent"
	"github.com/ntc-goer/microservice-examples/kitchen/ent/ticket"
)

type TicketRepo struct {
	TicketClient *ent.TicketClient
}

func NewTicketRepo(dc *ent.TicketClient) *TicketRepo {
	return &TicketRepo{
		TicketClient: dc,
	}
}

func (tr *TicketRepo) CreatePendingTicket(ctx context.Context, orderId, requestId string) (*ent.Ticket, error) {
	ticket, err := tr.TicketClient.Create().SetOrderID(orderId).SetRequestID(requestId).SetStatus(ticket.StatusCREATE_PENDING).Save(ctx)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (tr *TicketRepo) CancelTicket(ctx context.Context, ticketId string) (*ent.Ticket, error) {
	ticketUUID, err := uuid.Parse(ticketId)
	if err != nil {
		return nil, err
	}
	ticket, err := tr.TicketClient.UpdateOneID(ticketUUID).SetStatus(ticket.StatusCANCELED).Save(ctx)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (tr *TicketRepo) AcceptTicket(ctx context.Context, ticketId string) (*ent.Ticket, error) {
	ticketUUID, err := uuid.Parse(ticketId)
	if err != nil {
		return nil, err
	}
	ticket, err := tr.TicketClient.UpdateOneID(ticketUUID).SetStatus(ticket.StatusAWAITING_ACCEPTANCE).Save(ctx)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}
