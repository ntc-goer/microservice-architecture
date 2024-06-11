package inmem

import (
	"context"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
)

type Registry struct{}

func NewRegistry() (*Registry, error) {
	return &Registry{}, nil
}

func (reg *Registry) RegisterService(instanceId string, srvName string, srvAddr string, srvPort string, httpCheckUrl string) error {

	return nil
}

func (reg *Registry) Deregister(ctx context.Context, instanceID string) error {
	return nil
}

func (reg *Registry) HealthCheck(instanceID string) error {
	return nil
}

func (reg *Registry) Discover(ctx context.Context, serviceName string) ([]common.ActiveService, error) {
	return nil, nil
}
