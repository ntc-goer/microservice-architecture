//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/accounting/config"
	"github.com/ntc-goer/microservice-examples/accounting/service"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

type CoreDependency struct {
	Config           *config.Config
	Service          *service.CoreService
	ServiceDiscovery common.DiscoveryI
}

func NewCoreDependency(cfg *config.Config, coreSrv *service.CoreService, srvDis common.DiscoveryI) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		Service:          coreSrv,
		ServiceDiscovery: srvDis,
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
		NewCoreDependency)
	return &CoreDependency{}, nil
}
