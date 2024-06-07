//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/orderservice/service"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery/common"
	"github.com/ntc-goer/microservice-examples/registry/servicediscovery/consul"
)

type CoreDependency struct {
	Config           *config.Config
	ServiceImpl      *service.ServiceImpl
	ServiceDiscovery common.DiscoveryI
}

func NewCoreDependency(cfg *config.Config, srvImpl *service.ServiceImpl, srvDis common.DiscoveryI) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		ServiceImpl:      srvImpl,
		ServiceDiscovery: srvDis,
	}
}

//go:generate wire
func InitializeDependency(dcType string) (*CoreDependency, error) {
	wire.Build(
		config.Load,
		service.NewServiceImpl,
		//wire.Bind(new(common.DiscoveryI), new(*inmem.Registry)),
		//inmem.NewRegistry
		wire.Bind(new(common.DiscoveryI), new(*consul.Registry)),
		consul.NewRegistry,
		repository.NewRepository,
		NewCoreDependency)
	return &CoreDependency{}, nil
}
