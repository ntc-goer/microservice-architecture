package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/ent"
	"github.com/ntc-goer/microservice-examples/orderservice/ent/order"
)

type OrderRepo struct {
	OrderClient *ent.OrderClient
}

func NewOrderRepo(oc *ent.OrderClient) *OrderRepo {
	return &OrderRepo{
		OrderClient: oc,
	}
}
func (r *OrderRepo) CreatePendingOrder(ctx context.Context, orderId uuid.UUID, requestId uuid.UUID, addr string, userId string) (*ent.Order, error) {
	ord, err := r.OrderClient.
		Create().
		SetID(orderId).
		SetRequestID(requestId).SetStatus(order.StatusAPPROVAL_PENDING).SetAddress(addr).SetUserID(userId).Save(ctx)
	if err != nil {
		return nil, err
	}
	return ord, nil
}
