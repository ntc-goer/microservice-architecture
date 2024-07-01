package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/ent"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/ntc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type DishService struct {
	orderpb.UnimplementedDishServiceServer
	Repo  *repository.Repository
	Trace trace.Tracer
}

func NewDishService(repo *repository.Repository) *DishService {
	return &DishService{
		Repo:  repo,
		Trace: otel.Tracer("DishService"),
	}

}

func (s *DishService) GetOrderDish(ctx context.Context, req *orderpb.GetOrderDishRequest) (*orderpb.GetOrderDishResponse, error) {
	ctx, span := s.Trace.Start(ctx, "DishService.GetOrderDish")
	defer span.End()
	// Add attributes to the span
	span.SetAttributes(
		attribute.String("OrderId", req.OrderId))

	orderId, err := uuid.Parse(req.OrderId)
	if err != nil {
		return nil, err
	}
	dishes, err := s.Repo.Dish.GetDishesByOrderId(ctx, orderId)
	if err != nil {
		return nil, err
	}
	return &orderpb.GetOrderDishResponse{
		Dishes: ntc.Map(dishes, func(d *ent.Dish) *orderpb.DishItem {
			return &orderpb.DishItem{
				DishId: d.DishID,
				Dish:   d.DishName,
				Total:  int32(d.Quantity),
			}
		}),
	}, nil
}
