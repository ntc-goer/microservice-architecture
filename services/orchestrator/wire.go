//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/pkg"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/orderservice/service"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

type CoreDependency struct {
	Config           *config.Config
	ServiceImpl      *service.Impl
	ServiceDiscovery common.DiscoveryI
	DB *pkg.DB
}

func NewCoreDependency(cfg *config.Config, srvImpl *service.Impl, srvDis common.DiscoveryI, db *pkg.DB) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		ServiceImpl:      srvImpl,
		ServiceDiscovery: srvDis,
		DB: db,
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
		pkg.NewDB,
		pkg.NewLB,
		NewCoreDependency)
	return &CoreDependency{}, nil
}
