package service

import (
	"github.com/google/wire"
)

type CoreService struct {
	Health  *HealthService
	Kitchen *KitchenService
}

func NewCoreService(healthService *HealthService, kitchenService *KitchenService) *CoreService {
	return &CoreService{
		Health:  healthService,
		Kitchen: kitchenService,
	}
}

var WireSet = wire.NewSet(NewHealthService, NewKitchenService, NewCoreService)
