package consul

import (
	"context"
	consultapi "github.com/hashicorp/consul/api"
	"github.com/ntc-goer/microservice-examples/registry/serviceregistration/common"
	"log"
	"strconv"
)

type Registry struct {
	Client *consultapi.Client
}

func NewRegistry() (*Registry, error) {
	cfg := consultapi.DefaultConfig()
	client, err := consultapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Registry{
		Client: client,
	}, nil
}

func (reg *Registry) RegisterService(instanceId string, srvName string, srvAddr string, srvPort string, httpCheckUrl string) error {
	portInt, err := strconv.Atoi(srvPort)
	if err != nil {
		return err
	}
	registration := &consultapi.AgentServiceRegistration{
		ID:      instanceId,
		Name:    srvName,
		Address: srvAddr,
		Port:    portInt,
		Check: &consultapi.AgentServiceCheck{
			HTTP: httpCheckUrl,
			//CheckID: instanceId,
			//TLSSkipVerify: true,
			//TTL:                            "30s",
			Timeout:                        "1s",
			Interval:                       "10s",
			DeregisterCriticalServiceAfter: "60s",
		},
	}
	err = reg.Client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Service Registration fail")
	}
	return nil
}

func (reg *Registry) Deregister(ctx context.Context, instanceID string) error {
	log.Printf("Deregistering service %s", instanceID)
	return reg.Client.Agent().CheckDeregister(instanceID)
}

func (reg *Registry) HealthCheck(instanceID string) error {
	return reg.Client.Agent().UpdateTTL(instanceID, "online", consultapi.HealthPassing)
}

func (reg *Registry) Discover(ctx context.Context, serviceName string) ([]common.ActiveService, error) {
	entries, _, err := reg.Client.Health().Service(serviceName, "", true, nil)
	if err != nil {
		return nil, err
	}

	var instances []common.ActiveService
	for _, entry := range entries {
		instances = append(instances, common.ActiveService{
			ID:      entry.Service.ID,
			Service: entry.Service.Service,
			Address: entry.Service.Address,
			Port:    entry.Service.Port,
		})
	}

	return instances, nil
}
