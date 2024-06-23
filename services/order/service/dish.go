package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/ent"
	orderproto "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/ntc"
)

func (s *Impl) GetOrderDish(ctx context.Context, req *orderproto.GetOrderDishRequest) (*orderproto.GetOrderDishResponse, error) {
	orderId, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}
	dishes, err := s.Repo.Dish.GetDishesByOrderId(ctx, orderId)
	if err != nil {
		return nil, err
	}
	return &orderproto.GetOrderDishResponse{
		Dishes: ntc.Map(dishes, func(d *ent.Dish) *orderproto.OrderItem {
			return &orderproto.OrderItem{
				DishId: d.DishID,
				Dish:   d.DishName,
				Total:  int32(d.Quantity),
			}
		}),
	}, nil
}
