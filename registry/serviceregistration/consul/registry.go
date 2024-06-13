package consul

import (
	"context"
	"fmt"
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
	cfg.Address = "localhost:8500"
	client, err := consultapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &Registry{
		Client: client,
	}, nil
}

func (reg *Registry) RegisterService(instanceId string, srvName string, srvAddr string, srvPort string, checkType common.CheckType) error {
	portInt, err := strconv.Atoi(srvPort)
	agentServiceCheck := &consultapi.AgentServiceCheck{
		Interval:                       "3s",
		DeregisterCriticalServiceAfter: "3s",
	}
	if checkType == common.HTTP_CHECK_TYPE {
		agentServiceCheck.HTTP = fmt.Sprintf("http://%s:%s/health", srvAddr, srvPort)
	} else {
		agentServiceCheck.GRPC = fmt.Sprintf("%s:%s", srvAddr, srvPort)
	}
	if err != nil {
		return err
	}
	registration := &consultapi.AgentServiceRegistration{
		ID:      instanceId,
		Tags:    []string{fmt.Sprintf("urlprefix-/%s proto=grpc", srvName)},
		Name:    srvName,
		Address: srvAddr,
		Port:    portInt,
		Check:   agentServiceCheck,
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
