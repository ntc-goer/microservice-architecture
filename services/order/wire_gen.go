// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/pkg"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/orderservice/service"
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
	repositoryRepository, err := repository.NewRepository(configConfig)
	if err != nil {
		return nil, err
	}
	lb := pkg.NewLB(configConfig)
	msgQueue := queue.NewMsgQueue()
	impl, err := service.NewServiceImpl(registry, repositoryRepository, configConfig, lb, msgQueue)
	if err != nil {
		return nil, err
	}
	coreDependency := NewCoreDependency(configConfig, impl, registry, repositoryRepository)
	return coreDependency, nil
}

// wire.go:

type CoreDependency struct {
	Config           *config.Config
	ServiceImpl      *service.Impl
	ServiceDiscovery common.DiscoveryI
	Repository       *repository.Repository
}

func NewCoreDependency(cfg *config.Config, srvImpl *service.Impl, srvDis common.DiscoveryI, r *repository.Repository) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		ServiceImpl:      srvImpl,
		ServiceDiscovery: srvDis,
		Repository:       r,
	}
}
