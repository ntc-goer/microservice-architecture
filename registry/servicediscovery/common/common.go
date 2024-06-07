package common

import "context"

type DiscoveryType string

const (
	DC_CONSUL DiscoveryType = "consul"
	DC_INMEM  DiscoveryType = "inmem"
)

type ActiveService struct {
	ID      string
	Service string
	Address string
	Port    int
}

type DiscoveryI interface {
	RegisterService(instanceId string, srvName string, srvAddr string, srvPort string) error
	Deregister(ctx context.Context, instanceID string) error
	HealthCheck(instanceID string) error
	Discover(ctx context.Context, serviceName string) ([]ActiveService, error)
}
