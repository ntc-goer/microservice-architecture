package pkg

import (
	"fmt"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LB struct {
	Config *config.Config
}

func NewLB(cfg *config.Config) *LB {
	return &LB{
		Config: cfg,
	}
}

const _GRPC_CONFIG = `{
			"loadBalancingPolicy": "round_robin", 
			"healthCheckConfig": {"serviceName": "%s"}
            "methodConfig": [{
                "name": [{"service": "%s"}],
                "waitForReady": true,
                "retryPolicy": {
                    "MaxAttempts": 2,
                    "InitialBackoff": "1s",
                    "MaxBackoff": "5s",
                    "BackoffMultiplier": 1.0,
                    "RetryableStatusCodes": [ "UNAVAILABLE" ]
                }
            }]
        }`

func (lb *LB) GetConnection(srvName string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		lb.Config.LBServiceHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(
			fmt.Sprintf(_GRPC_CONFIG, srvName, srvName)))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
