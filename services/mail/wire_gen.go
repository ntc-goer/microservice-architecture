// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ntc-goer/microservice-examples/mailservice/config"
	"github.com/ntc-goer/microservice-examples/mailservice/service"
	"github.com/ntc-goer/microservice-examples/registry/queue"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

// Injectors from wire.go:

//go:generate wire
func InitializeDependency(dcType string) (*CoreDependency, error) {
	configConfig, err := config.Load()
	if err != nil {
		return nil, err
	}
	registry, err := consul.NewRegistry()
	if err != nil {
		return nil, err
	}
	msgQueue := queue.NewMsgQueue()
	serviceService := service.NewService(configConfig)
	coreDependency := NewCoreDependency(configConfig, registry, msgQueue, serviceService)
	return coreDependency, nil
}

// wire.go:

type CoreDependency struct {
	Config           *config.Config
	ServiceDiscovery common.DiscoveryI
	Queue            *queue.MsgQueue
	Service          *service.Service
}

func NewCoreDependency(cfg *config.Config, srvDis common.DiscoveryI, queue2 *queue.MsgQueue, srv *service.Service) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		ServiceDiscovery: srvDis,
		Queue:            queue2,
		Service:          srv,
	}
}
