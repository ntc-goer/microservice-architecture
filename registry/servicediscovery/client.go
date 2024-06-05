package servicediscovery

import (
	consultapi "github.com/hashicorp/consul/api"
	"log"
	"strconv"
	"time"
)

type ServiceDiscovery struct {
	Client *consultapi.Client
}

const ttl = time.Second * 8

func NewServiceDiscovery() (*ServiceDiscovery, error) {
	cfg := consultapi.DefaultConfig()
	client, err := consultapi.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return &ServiceDiscovery{
		Client: client,
	}, nil
}

func (sd *ServiceDiscovery) RegisterService(srvId string, srvName string, srvAddr string, srvPort string) error {
	portInt, err := strconv.Atoi(srvPort)
	if err != nil {
		return err
	}
	registration := &consultapi.AgentServiceRegistration{
		ID:      srvId,
		Name:    srvName,
		Address: srvAddr,
		Port:    portInt,
		Check: &consultapi.AgentServiceCheck{
			GRPC:                           "127.0.0.1:50000",
			Timeout:                        "5s",
			Interval:                       "5s",
			DeregisterCriticalServiceAfter: "10s",
		},
	}
	err = sd.Client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Service Registration fail")
	}
	return nil
}
