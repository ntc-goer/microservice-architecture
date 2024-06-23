package service

import (
	"github.com/google/wire"
)

type CoreService struct {
	Health *HealthService
	User   *UserService
}

func NewCoreService(healthService *HealthService, userService *UserService) *CoreService {
	return &CoreService{
		Health: healthService,
		User:   userService,
	}
}

var WireSet = wire.NewSet(NewHealthService, NewUserService, NewCoreService)
