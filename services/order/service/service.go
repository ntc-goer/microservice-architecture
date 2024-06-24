package service

import (
	"github.com/google/wire"
)

type CoreService struct {
	Health *HealthService
	Order  *OrderService
	Dish   *DishService
}

func NewCoreService(healthService *HealthService, orderService *OrderService, dishService *DishService) *CoreService {
	return &CoreService{
		Order:  orderService,
		Health: healthService,
		Dish:   dishService,
	}
}

var WireSet = wire.NewSet(NewHealthService, NewOrderService, NewDishService, NewCoreService)
