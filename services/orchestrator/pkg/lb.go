package pkg

import (
	"github.com/ntc-goer/microservice-examples/orchestrator/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

type LB struct {
	Config *config.Config
}

const _GRPC_CONFIG = `{
			"loadBalancingPolicy": "round_robin", 
			"healthCheckConfig": {"serviceName": "%s"},
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

func GetConnection(addr string, srvName string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getEnv(key string, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	return val
}

func GetGRPCClient[T any](lbHost string, srvName string, f func(grpc.ClientConnInterface) T) (T, error) {
	appEnv := getEnv("APP_ENV", "local")
	if appEnv != "local" {
		lbHost = srvName
	}

	conn, err := GetConnection(lbHost, srvName)
	if err != nil {
		var zero T
		return zero, err
	}
	client := f(conn)
	return client, nil
}
