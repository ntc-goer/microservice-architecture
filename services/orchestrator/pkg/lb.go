package pkg

import (
	"fmt"
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LB struct {
	Config *config.Config
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

func GetConnection(lbHost string, srvName string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		lbHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(
			fmt.Sprintf(_GRPC_CONFIG, srvName, srvName)))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func GetGRPCClient[T any](lbHost string, srvName string, f func(grpc.ClientConnInterface) T) (T, error) {
	conn, err := GetConnection(lbHost, srvName)
	if err != nil {
		var zero T
		return zero, err
	}
	client := f(conn)
	return client, nil
}
