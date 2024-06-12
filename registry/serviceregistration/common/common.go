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

type CheckType string

const (
	HTTP_CHECK_TYPE CheckType = "http"
	GRPC_CHECK_TYPE CheckType = "grpc"
)

type DiscoveryI interface {
	RegisterService(instanceId string, srvName string, srvAddr string, srvPort string, checkType CheckType) error
	Deregister(ctx context.Context, instanceID string) error
	HealthCheck(instanceID string) error
	Discover(ctx context.Context, serviceName string) ([]ActiveService, error)
}
