package main

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/kitchen/config"
	"github.com/ntc-goer/microservice-examples/kitchen/repository"
	"github.com/ntc-goer/microservice-examples/kitchen/service"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

type CoreDependency struct {
	Config           *config.Config
	Service          *service.CoreService
	ServiceDiscovery common.DiscoveryI
	Repository *repository.Repository
}

func NewCoreDependency(cfg *config.Config, coreSrv *service.CoreService, srvDis common.DiscoveryI, repo *repository.Repository) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		Service:          coreSrv,
		ServiceDiscovery: srvDis,
		Repository: repo,
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
		NewCoreDependency)
	return &CoreDependency{}, nil
}
