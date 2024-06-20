//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
	"github.com/ntc-goer/microservice-examples/orchestrator/service"
	"github.com/ntc-goer/microservice-examples/registry/broker"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

type CoreDependency struct {
	Config           *config.Config
	CoreService      *service.CoreService
	ServiceDiscovery common.DiscoveryI
	Broker           *broker.Broker
}

func NewCoreDependency(cfg *config.Config, coreSrv *service.CoreService, srvDis common.DiscoveryI, br *broker.Broker) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		ServiceDiscovery: srvDis,
		CoreService:      coreSrv,
		Broker:           br,
	}
}

//go:generate wire
func InitializeDependency(dcType string) (*CoreDependency, error) {
	wire.Build(
		config.Load,
		//wire.Bind(new(common.DiscoveryI), new(*inmem.Registry)),
		//inmem.NewRegistry
		wire.Bind(new(common.DiscoveryI), new(*consul.Registry)),
		consul.NewRegistry,
		//pkg.NewLB,
		service.WireSet,
		broker.NewBroker,
		NewCoreDependency)
	return &CoreDependency{}, nil
}
