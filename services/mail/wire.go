//go:build wireinject
// +build wireinject
package main
import (
	"github.com/google/wire"
	"github.com/ntc-goer/microservice-examples/mailservice/config"
	"github.com/ntc-goer/microservice-examples/mailservice/service"
	"github.com/ntc-goer/microservice-examples/registry/queue"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/consul"
)

type CoreDependency struct {
	Config           *config.Config
	ServiceDiscovery common.DiscoveryI
	Queue *queue.MsgQueue
	Service *service.Service
}

func NewCoreDependency(cfg *config.Config, srvDis common.DiscoveryI, queue *queue.MsgQueue, srv *service.Service) *CoreDependency {
	return &CoreDependency{
		Config:           cfg,
		ServiceDiscovery: srvDis,
		Queue: queue,
		Service: srv,
	}
}

//go:generate wire
func InitializeDependency(dcType string) (*CoreDependency, error) {
	wire.Build(
		config.Load,
		//wire.Bind(new(common.DiscoveryI), new(*inmem.Registry)),
		//inmem.NewRegistry
		wire.Bind(new(common.DiscoveryI), new(*consul.Registry)),
		service.NewService,
		consul.NewRegistry,
		queue.NewMsgQueue,
		NewCoreDependency)
	return &CoreDependency{}, nil
}
