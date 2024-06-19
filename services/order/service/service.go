package service

import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/pkg"
	orderpb "github.com/ntc-goer/microservice-examples/orderservice/proto"
	"github.com/ntc-goer/microservice-examples/orderservice/repository"
	"github.com/ntc-goer/microservice-examples/registry/broker"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
)

type Impl struct {
	orderpb.UnimplementedOrderServiceServer
	SrvDis      common.DiscoveryI
	Repo        *repository.Repository
	Config      *config.Config
	LoadBalance *pkg.LB
	Queue       *broker.Broker
}

func NewServiceImpl(srvDis common.DiscoveryI, repo *repository.Repository, cfg *config.Config, lb *pkg.LB, q *broker.Broker) (*Impl, error) {
	return &Impl{
		SrvDis:      srvDis,
		Repo:        repo,
		Config:      cfg,
		LoadBalance: lb,
		Queue:       q,
	}, nil
}

type CoreService struct {
	Health      *HealthService
	ServiceImpl *Impl
}

func NewCoreService(healthService *HealthService, ServiceImpl *Impl) *CoreService {
	return &CoreService{
		ServiceImpl: ServiceImpl,
		Health:      healthService,
	}
}

var WireSet = wire.NewSet(NewHealthService, NewServiceImpl, NewCoreService)
