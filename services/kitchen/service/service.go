package service

import (
	"github.com/google/wire"
)

type CoreService struct {
	Health     *HealthService
	Accounting *AccountingService
}

func NewCoreService(healthService *HealthService, accountingService *AccountingService) *CoreService {
	return &CoreService{
		Health:     healthService,
		Accounting: accountingService,
	}
}

var WireSet = wire.NewSet(NewHealthService, NewKitchenService, NewCoreService)
