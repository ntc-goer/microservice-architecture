//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/pkg"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/orderservice/service"
	"github.com/ntc-goer/microservice-examples/registry/broker"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

type CoreDependency struct {
	Config           *config.Config
	CoreService      *service.CoreService
	ServiceDiscovery common.DiscoveryI
	Repository       *repository.Repository
	Broker *broker.Broker
}

func NewCoreDependency(cfg *config.Config, coreSrv *service.CoreService, srvDis common.DiscoveryI, r *repository.Repository, br *broker.Broker) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		CoreService: coreSrv,
		ServiceDiscovery: srvDis,
		Repository:       r,
		Broker: br,
	}
}

//go:generate wire
func InitializeDependency(dcType string) (*CoreDependency, error) {
	wire.Build(
		config.Load,
		service.WireSet,
		//wire.Bind(new(common.DiscoveryI), new(*inmem.Registry)),
		//inmem.NewRegistry
		wire.Bind(new(common.DiscoveryI), new(*consul.Registry)),
		consul.NewRegistry,
		repository.NewRepository,
		broker.NewBroker,
		pkg.NewLB,
		NewCoreDependency)
	return &CoreDependency{}, nil
}
